#!/bin/bash
#
# NATS Cluster Health Check
# Comprehensive health monitoring for NATS cluster
#
# Usage: ./NATS_health-check.sh [--verbose]
# Exit codes:
#   0 - All healthy
#   1 - Cluster formation issues
#   2 - JetStream issues
#   3 - Stream/Consumer issues
#   4 - Node health issues
#

set -euo pipefail

# Configuration
NATS_NODES=("nats-1:4222" "nats-2:4222" "nats-3:4222")
NATS_CONTAINERS=("nats-1" "nats-2" "nats-3")
VERBOSE=false

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

# Parse arguments
while [[ $# -gt 0 ]]; do
    case $1 in
        --verbose|-v)
            VERBOSE=true
            shift
            ;;
        --help|-h)
            echo "NATS Cluster Health Check"
            echo "Usage: $0 [--verbose]"
            echo "Exit codes: 0=healthy, 1=cluster, 2=jetstream, 3=streams, 4=nodes"
            exit 0
            ;;
        *)
            echo "Unknown option: $1"
            exit 1
            ;;
    esac
done

# Logging functions
log() {
    if [[ "$VERBOSE" == "true" ]]; then
        echo -e "${BLUE}[$(date '+%H:%M:%S')] $1${NC}"
    fi
}

log_success() {
    echo -e "${GREEN}✓ $1${NC}"
}

log_error() {
    echo -e "${RED}✗ $1${NC}"
}

log_warning() {
    echo -e "${YELLOW}⚠ $1${NC}"
}

# Health check functions
check_node_health() {
    local node="$1"
    local container_name=$(echo "$node" | cut -d':' -f1)
    
    log "Checking node health: $node"
    
    # Check if container is running
    if ! docker ps --format '{{.Names}}' | grep -q "^${container_name}$"; then
        log_error "Container $container_name is not running"
        return 1
    fi
    
    # Check HTTP endpoint
    if ! curl -sf "http://$node/healthz" >/dev/null 2>&1; then
        log_error "Health endpoint not responding: $node"
        return 1
    fi
    
    # Check NATS connectivity
    if ! nats server check --server="$node" >/dev/null 2>&1; then
        log_error "NATS server not responding: $node"
        return 1
    fi
    
    log_success "Node $node is healthy"
    return 0
}

check_cluster_formation() {
    log "Checking cluster formation..."
    
    local leader_count=0
    local follower_count=0
    local total_nodes=0
    local cluster_size=0
    
    for node in "${NATS_NODES[@]}"; do
        if ! curl -sf "http://$node/varz" >/dev/null 2>&1; then
            continue
        fi
        
        local varz
        varz=$(curl -s "http://$node/varz" 2>/dev/null || echo "{}")
        
        if echo "$varz" | jq -e '.cluster' >/dev/null 2>&1; then
            total_nodes=$((total_nodes + 1))
            
            if echo "$varz" | jq -r '.cluster.leader' 2>/dev/null | grep -q "true"; then
                leader_count=$((leader_count + 1))
            else
                follower_count=$((follower_count + 1))
            fi
            
            # Get cluster size from first responding node
            if [[ $cluster_size -eq 0 ]]; then
                cluster_size=$(echo "$varz" | jq -r '.cluster.urls // [] | length' 2>/dev/null || echo 0)
                cluster_size=$((cluster_size + 1)) # Add self
            fi
        fi
    done
    
    # Validate cluster state
    local healthy=true
    
    if [[ $leader_count -ne 1 ]]; then
        log_error "Expected 1 leader, found $leader_count"
        healthy=false
    fi
    
    if [[ $total_nodes -lt 2 ]]; then
        log_error "Insufficient nodes online: $total_nodes (minimum 2 for cluster)"
        healthy=false
    fi
    
    if [[ $cluster_size -ne ${#NATS_NODES[@]} ]]; then
        log_warning "Cluster size mismatch: expected ${#NATS_NODES[@]}, got $cluster_size"
    fi
    
    if [[ "$healthy" == "true" ]]; then
        log_success "Cluster formation healthy: 1 leader, $follower_count followers, $total_nodes total nodes"
        return 0
    else
        log_error "Cluster formation unhealthy"
        return 1
    fi
}

check_jetstream_status() {
    log "Checking JetStream status..."
    
    local js_enabled_count=0
    local js_leader_found=false
    
    for node in "${NATS_NODES[@]}"; do
        if ! nats server check --server="$node" >/dev/null 2>&1; then
            continue
        fi
        
        # Check if JetStream is enabled
        if nats server info --server="$node" 2>/dev/null | grep -q "JetStream: Enabled"; then
            js_enabled_count=$((js_enabled_count + 1))
            
            # Check if this node is JetStream cluster leader
            local jsz
            jsz=$(curl -s "http://$node/jsz" 2>/dev/null || echo "{}")
            if echo "$jsz" | jq -r '.meta.leader' 2>/dev/null | grep -q "true"; then
                js_leader_found=true
            fi
        fi
    done
    
    if [[ $js_enabled_count -eq 0 ]]; then
        log_error "JetStream not enabled on any node"
        return 1
    fi
    
    if [[ "$js_leader_found" != "true" ]]; then
        log_error "No JetStream meta leader found"
        return 1
    fi
    
    log_success "JetStream healthy: enabled on $js_enabled_count nodes with meta leader"
    return 0
}

check_streams_and_consumers() {
    log "Checking streams and consumers..."
    
    # Find available server
    local server=""
    for node in "${NATS_NODES[@]}"; do
        if nats server check --server="$node" >/dev/null 2>&1; then
            server="$node"
            break
        fi
    done
    
    if [[ -z "$server" ]]; then
        log_error "No available NATS server for stream check"
        return 1
    fi
    
    # Get stream list
    local streams
    streams=$(nats stream list --server="$server" 2>/dev/null || echo "")
    
    if [[ -z "$streams" ]]; then
        log_warning "No streams found (this may be normal for a fresh cluster)"
        return 0
    fi
    
    local unhealthy_streams=0
    local total_streams=0
    
    while IFS= read -r stream; do
        if [[ -z "$stream" ]]; then continue; fi
        
        total_streams=$((total_streams + 1))
        log "Checking stream: $stream"
        
        # Get stream info
        local stream_info
        stream_info=$(nats stream info "$stream" --server="$server" --json 2>/dev/null || echo "{}")
        
        if [[ -z "$stream_info" || "$stream_info" == "{}" ]]; then
            log_error "Could not get info for stream: $stream"
            unhealthy_streams=$((unhealthy_streams + 1))
            continue
        fi
        
        # Check stream state
        local state
        state=$(echo "$stream_info" | jq -r '.state.state // "unknown"' 2>/dev/null)
        
        if [[ "$state" != "STREAM_STATE_AVAILABLE" && "$state" != "available" ]]; then
            log_error "Stream $stream is in state: $state"
            unhealthy_streams=$((unhealthy_streams + 1))
            continue
        fi
        
        # Check replicas if configured
        local replicas
        replicas=$(echo "$stream_info" | jq -r '.config.num_replicas // 1' 2>/dev/null)
        
        if [[ $replicas -gt 1 ]]; then
            local replica_info
            replica_info=$(echo "$stream_info" | jq -r '.cluster.replicas // []' 2>/dev/null)
            local healthy_replicas
            healthy_replicas=$(echo "$replica_info" | jq '[.[] | select(.current == true)] | length' 2>/dev/null || echo 0)
            
            if [[ $healthy_replicas -lt $replicas ]]; then
                log_error "Stream $stream: only $healthy_replicas/$replicas replicas are current"
                unhealthy_streams=$((unhealthy_streams + 1))
                continue
            fi
        fi
        
        log_success "Stream $stream is healthy"
        
    done <<< "$streams"
    
    if [[ $unhealthy_streams -eq 0 ]]; then
        if [[ $total_streams -gt 0 ]]; then
            log_success "All $total_streams streams are healthy"
        else
            log_success "No streams to check"
        fi
        return 0
    else
        log_error "$unhealthy_streams/$total_streams streams are unhealthy"
        return 1
    fi
}

get_cluster_summary() {
    echo
    echo "=== NATS CLUSTER HEALTH SUMMARY ==="
    
    local leader=""
    local online_nodes=0
    
    for node in "${NATS_NODES[@]}"; do
        local status="OFFLINE"
        local role="Unknown"
        
        if curl -sf "http://$node/healthz" >/dev/null 2>&1; then
            status="ONLINE"
            online_nodes=$((online_nodes + 1))
            
            local varz
            varz=$(curl -s "http://$node/varz" 2>/dev/null || echo "{}")
            
            if echo "$varz" | jq -r '.cluster.leader' 2>/dev/null | grep -q "true"; then
                role="LEADER"
                leader="$node"
            else
                role="FOLLOWER"
            fi
        fi
        
        printf "%-15s %-8s %-10s\n" "$node" "$status" "$role"
    done
    
    echo
    echo "Online Nodes: $online_nodes/${#NATS_NODES[@]}"
    if [[ -n "$leader" ]]; then
        echo "Cluster Leader: $leader"
    else
        echo "Cluster Leader: NONE (unhealthy)"
    fi
}

# Main health check execution
main() {
    local exit_code=0
    
    echo "NATS Cluster Health Check"
    echo "========================="
    
    # 1. Check individual node health
    log "Phase 1: Checking individual node health..."
    local unhealthy_nodes=0
    
    for node in "${NATS_NODES[@]}"; do
        if ! check_node_health "$node"; then
            unhealthy_nodes=$((unhealthy_nodes + 1))
        fi
    done
    
    if [[ $unhealthy_nodes -gt 0 ]]; then
        log_error "Node health issues detected: $unhealthy_nodes unhealthy nodes"
        exit_code=4
    fi
    
    # 2. Check cluster formation
    log "Phase 2: Checking cluster formation..."
    if ! check_cluster_formation; then
        log_error "Cluster formation issues detected"
        if [[ $exit_code -eq 0 ]]; then exit_code=1; fi
    fi
    
    # 3. Check JetStream status
    log "Phase 3: Checking JetStream status..."
    if ! check_jetstream_status; then
        log_error "JetStream issues detected"
        if [[ $exit_code -eq 0 ]]; then exit_code=2; fi
    fi
    
    # 4. Check streams and consumers
    log "Phase 4: Checking streams and consumers..."
    if ! check_streams_and_consumers; then
        log_error "Stream/Consumer issues detected"
        if [[ $exit_code -eq 0 ]]; then exit_code=3; fi
    fi
    
    # Show summary
    get_cluster_summary
    
    echo
    if [[ $exit_code -eq 0 ]]; then
        log_success "ALL HEALTH CHECKS PASSED! Cluster is healthy."
    else
        log_error "HEALTH CHECK FAILED! Exit code: $exit_code"
        case $exit_code in
            1) echo "Issue: Cluster formation problems" ;;
            2) echo "Issue: JetStream problems" ;;
            3) echo "Issue: Stream/Consumer problems" ;;
            4) echo "Issue: Node health problems" ;;
        esac
    fi
    
    exit $exit_code
}

# Execute main function
main "$@"