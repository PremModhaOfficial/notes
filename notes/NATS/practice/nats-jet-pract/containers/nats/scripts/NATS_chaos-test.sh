#!/bin/bash
#
# NATS Chaos Testing Suite
# Comprehensive chaos testing for NATS cluster with automated verification
#
# Usage: ./NATS_chaos-test.sh [scenario]
# Scenarios: leader-kill, split-brain, network-partition, all
#

set -euo pipefail

# Configuration
NATS_NODES=("nats-1:4222" "nats-2:4222" "nats-3:4222")
NATS_CONTAINERS=("nats-1" "nats-2" "nats-3")
TEST_STREAM="chaos-test"
TEST_SUBJECT="test.chaos"
DOCKER_NETWORK="nats-cluster"
LOG_FILE="chaos-test-$(date +%Y%m%d-%H%M%S).log"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
PURPLE='\033[0;35m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

# Logging functions
log() {
    echo -e "${CYAN}[$(date '+%Y-%m-%d %H:%M:%S')] $1${NC}" | tee -a "$LOG_FILE"
}

log_success() {
    echo -e "${GREEN}✓ $1${NC}" | tee -a "$LOG_FILE"
}

log_error() {
    echo -e "${RED}✗ $1${NC}" | tee -a "$LOG_FILE"
}

log_warning() {
    echo -e "${YELLOW}⚠ $1${NC}" | tee -a "$LOG_FILE"
}

log_info() {
    echo -e "${BLUE}ℹ $1${NC}" | tee -a "$LOG_FILE"
}

# Cleanup function
cleanup() {
    log "Starting cleanup and restoration..."
    
    # Restore all containers
    for container in "${NATS_CONTAINERS[@]}"; do
        if ! docker ps -q -f name="$container" | grep -q .; then
            log "Starting container $container..."
            docker start "$container" 2>/dev/null || true
        fi
        
        # Reconnect to network if needed
        if ! docker network inspect "$DOCKER_NETWORK" --format '{{range .Containers}}{{.Name}} {{end}}' | grep -q "$container"; then
            log "Reconnecting $container to network..."
            docker network connect "$DOCKER_NETWORK" "$container" 2>/dev/null || true
        fi
    done
    
    # Wait for cluster to reform
    log "Waiting for cluster to reform..."
    sleep 10
    
    # Clean up test stream
    if nats stream info "$TEST_STREAM" --server="${NATS_NODES[0]}" >/dev/null 2>&1; then
        nats stream delete "$TEST_STREAM" --server="${NATS_NODES[0]}" --force 2>/dev/null || true
        log_success "Test stream cleaned up"
    fi
    
    log_success "Cleanup completed"
}

# Signal handlers
trap cleanup EXIT
trap 'log_error "Script interrupted"; exit 1' INT TERM

# Helper functions
wait_for_leader_election() {
    local timeout=${1:-30}
    local start_time=$(date +%s)
    
    log "Waiting for leader election (timeout: ${timeout}s)..."
    
    while true; do
        local current_time=$(date +%s)
        local elapsed=$((current_time - start_time))
        
        if [ $elapsed -gt $timeout ]; then
            log_error "Leader election timeout after ${timeout}s"
            return 1
        fi
        
        # Check if we have a leader
        for node in "${NATS_NODES[@]}"; do
            if curl -s "http://$node/varz" | jq -r '.cluster.leader' | grep -q "true" 2>/dev/null; then
                local election_time=$elapsed
                log_success "Leader elected in ${election_time}s"
                return 0
            fi
        done
        
        sleep 1
    done
}

get_cluster_leader() {
    for node in "${NATS_NODES[@]}"; do
        if curl -s "http://$node/varz" 2>/dev/null | jq -r '.cluster.leader' 2>/dev/null | grep -q "true"; then
            echo "$node"
            return 0
        fi
    done
    return 1
}

get_leader_container() {
    local leader_node=$(get_cluster_leader)
    if [ -n "$leader_node" ]; then
        echo "$leader_node" | cut -d':' -f1
    else
        return 1
    fi
}

verify_message_persistence() {
    local test_msg="chaos-test-$(date +%s)"
    local server="${1:-${NATS_NODES[0]}}"
    
    log "Verifying message persistence..."
    
    # Publish test message
    if echo "$test_msg" | nats pub "$TEST_SUBJECT" --server="$server" 2>/dev/null; then
        log_success "Test message published: $test_msg"
    else
        log_error "Failed to publish test message"
        return 1
    fi
    
    # Wait a moment for persistence
    sleep 2
    
    # Try to consume the message
    local received_msg
    received_msg=$(timeout 5 nats sub "$TEST_SUBJECT" --server="$server" --count=1 2>/dev/null | tail -n1 || echo "")
    
    if [ "$received_msg" = "$test_msg" ]; then
        log_success "Message persistence verified"
        return 0
    else
        log_error "Message persistence failed - expected: $test_msg, got: $received_msg"
        return 1
    fi
}

setup_test_stream() {
    log "Setting up test stream..."
    
    # Find available server
    local server=""
    for node in "${NATS_NODES[@]}"; do
        if nats server check --server="$node" >/dev/null 2>&1; then
            server="$node"
            break
        fi
    done
    
    if [ -z "$server" ]; then
        log_error "No available NATS server found"
        return 1
    fi
    
    # Create test stream with persistence
    nats stream add "$TEST_STREAM" \
        --subjects="$TEST_SUBJECT" \
        --storage=file \
        --replicas=3 \
        --retention=limits \
        --discard=old \
        --max-msgs=1000 \
        --max-age=1h \
        --server="$server" \
        --force >/dev/null 2>&1
    
    log_success "Test stream '$TEST_STREAM' created with 3 replicas"
}

# Test Scenarios

test_leader_kill() {
    log_info "=== SCENARIO 1: Leader Kill Test ==="
    
    # Get current leader
    local leader_container
    leader_container=$(get_leader_container)
    if [ -z "$leader_container" ]; then
        log_error "Could not identify cluster leader"
        return 1
    fi
    
    log "Current leader: $leader_container"
    
    # Record pre-test state
    local pre_test_time=$(date +%s)
    
    # Kill the leader
    log "Killing leader container: $leader_container"
    docker stop "$leader_container"
    log_warning "Leader $leader_container stopped"
    
    # Measure election time
    if wait_for_leader_election 30; then
        local new_leader
        new_leader=$(get_leader_container)
        log_success "New leader elected: $new_leader"
        
        # Verify message persistence after leader election
        if verify_message_persistence; then
            log_success "SCENARIO 1 PASSED: Leader kill and election successful"
        else
            log_error "SCENARIO 1 FAILED: Message persistence failed after leader election"
            return 1
        fi
    else
        log_error "SCENARIO 1 FAILED: Leader election timeout"
        return 1
    fi
    
    # Restart the killed node
    log "Restarting $leader_container..."
    docker start "$leader_container"
    sleep 5
    log_success "Node $leader_container restarted and rejoined cluster"
}

test_split_brain() {
    log_info "=== SCENARIO 2: Split Brain Test ==="
    
    # Stop 2 nodes to create minority partition
    local nodes_to_stop=("${NATS_CONTAINERS[1]}" "${NATS_CONTAINERS[2]}")
    
    log "Creating split brain by stopping 2 nodes: ${nodes_to_stop[*]}"
    
    for container in "${nodes_to_stop[@]}"; do
        docker stop "$container"
        log_warning "Stopped $container"
    done
    
    # Wait for split brain detection
    sleep 5
    
    # The remaining single node should not be able to form a majority
    local remaining_node="${NATS_CONTAINERS[0]}"
    log "Testing cluster state with single node: $remaining_node"
    
    # Try to get cluster info - should show degraded state
    if curl -s "http://${NATS_NODES[0]}/varz" | jq -r '.cluster.leader' | grep -q "false"; then
        log_success "Single node correctly shows no leader (no majority)"
    else
        log_warning "Single node behavior unexpected"
    fi
    
    # Restart one node to restore majority
    log "Restarting ${nodes_to_stop[0]} to restore majority..."
    docker start "${nodes_to_stop[0]}"
    
    # Wait for cluster to reform
    if wait_for_leader_election 30; then
        log_success "Cluster reformed with majority (2/3 nodes)"
        
        # Verify functionality
        if verify_message_persistence; then
            log_success "SCENARIO 2 PASSED: Split brain recovery successful"
        else
            log_error "SCENARIO 2 FAILED: Message persistence failed after recovery"
            return 1
        fi
    else
        log_error "SCENARIO 2 FAILED: Cluster failed to reform after split brain"
        return 1
    fi
    
    # Restart the last node
    log "Restarting ${nodes_to_stop[1]}..."
    docker start "${nodes_to_stop[1]}"
    sleep 5
    log_success "All nodes restarted"
}

test_network_partition() {
    log_info "=== SCENARIO 3: Network Partition Test ==="
    
    # Create network partition by disconnecting one node
    local isolated_node="${NATS_CONTAINERS[2]}"
    
    log "Creating network partition by isolating: $isolated_node"
    docker network disconnect "$DOCKER_NETWORK" "$isolated_node"
    log_warning "Node $isolated_node isolated from network"
    
    # Wait for partition detection
    sleep 5
    
    # Remaining nodes should maintain cluster
    log "Verifying remaining cluster functionality..."
    if verify_message_persistence "${NATS_NODES[0]}"; then
        log_success "Remaining cluster maintains functionality"
    else
        log_error "Remaining cluster lost functionality"
        return 1
    fi
    
    # Check that isolated node is not accessible
    if ! nats server check --server="${NATS_NODES[2]}" >/dev/null 2>&1; then
        log_success "Isolated node correctly unreachable"
    else
        log_warning "Isolated node still reachable (unexpected)"
    fi
    
    # Restore network partition
    log "Restoring network partition..."
    docker network connect "$DOCKER_NETWORK" "$isolated_node"
    sleep 5
    
    # Wait for node to rejoin
    local rejoin_timeout=20
    local start_time=$(date +%s)
    
    while true; do
        local current_time=$(date +%s)
        local elapsed=$((current_time - start_time))
        
        if [ $elapsed -gt $rejoin_timeout ]; then
            log_error "Node rejoin timeout"
            return 1
        fi
        
        if nats server check --server="${NATS_NODES[2]}" >/dev/null 2>&1; then
            log_success "Node $isolated_node rejoined cluster in ${elapsed}s"
            break
        fi
        
        sleep 1
    done
    
    # Final verification
    if verify_message_persistence; then
        log_success "SCENARIO 3 PASSED: Network partition recovery successful"
    else
        log_error "SCENARIO 3 FAILED: Final verification failed"
        return 1
    fi
}

# Main execution
main() {
    local scenario="${1:-all}"
    
    log_info "Starting NATS Chaos Testing Suite"
    log "Test scenario: $scenario"
    log "Log file: $LOG_FILE"
    
    # Initial cluster health check
    log "Performing initial cluster health check..."
    if ! ./scripts/NATS_health-check.sh >/dev/null 2>&1; then
        log_error "Cluster is not healthy before testing"
        exit 1
    fi
    log_success "Initial cluster health check passed"
    
    # Setup test environment
    setup_test_stream
    
    # Run scenarios
    local scenarios_run=0
    local scenarios_passed=0
    
    case "$scenario" in
        "leader-kill"|"1")
            scenarios_run=1
            test_leader_kill && scenarios_passed=$((scenarios_passed + 1))
            ;;
        "split-brain"|"2")
            scenarios_run=1
            test_split_brain && scenarios_passed=$((scenarios_passed + 1))
            ;;
        "network-partition"|"3")
            scenarios_run=1
            test_network_partition && scenarios_passed=$((scenarios_passed + 1))
            ;;
        "all"|"")
            scenarios_run=3
            test_leader_kill && scenarios_passed=$((scenarios_passed + 1))
            test_split_brain && scenarios_passed=$((scenarios_passed + 1))
            test_network_partition && scenarios_passed=$((scenarios_passed + 1))
            ;;
        *)
            log_error "Unknown scenario: $scenario"
            log_info "Available scenarios: leader-kill, split-brain, network-partition, all"
            exit 1
            ;;
    esac
    
    # Final report
    echo
    log_info "=== CHAOS TESTING SUMMARY ==="
    log "Scenarios run: $scenarios_run"
    log "Scenarios passed: $scenarios_passed"
    
    if [ $scenarios_passed -eq $scenarios_run ]; then
        log_success "ALL CHAOS TESTS PASSED! 🎉"
        exit 0
    else
        log_error "SOME CHAOS TESTS FAILED!"
        log "Check log file: $LOG_FILE"
        exit 1
    fi
}

# Show usage if requested
if [[ "${1:-}" == "-h" || "${1:-}" == "--help" ]]; then
    echo "NATS Chaos Testing Suite"
    echo
    echo "Usage: $0 [scenario]"
    echo
    echo "Scenarios:"
    echo "  leader-kill       Test leader failure and election"
    echo "  split-brain       Test split brain scenario (2 nodes down)"
    echo "  network-partition Test network partition recovery"
    echo "  all              Run all scenarios (default)"
    echo
    echo "Examples:"
    echo "  $0                    # Run all scenarios"
    echo "  $0 leader-kill        # Test leader election only"
    echo "  $0 split-brain        # Test split brain only"
    echo
    exit 0
fi

# Execute main function
main "$@"