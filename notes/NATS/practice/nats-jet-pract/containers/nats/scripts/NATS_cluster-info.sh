#!/bin/bash
#
# NATS Cluster Information Reporter
# Detailed cluster state and performance metrics
#
# Usage: ./NATS_cluster-info.sh [--json] [--metrics] [--streams]
#

set -euo pipefail

# Configuration
NATS_NODES=("nats-1:4222" "nats-2:4222" "nats-3:4222")
NATS_CONTAINERS=("nats-1" "nats-2" "nats-3")

# Output options
OUTPUT_JSON=false
SHOW_METRICS=false
SHOW_STREAMS=false
SHOW_ALL=true

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
BOLD='\033[1m'
NC='\033[0m'

# Parse arguments
while [[ $# -gt 0 ]]; do
    case $1 in
        --json)
            OUTPUT_JSON=true
            shift
            ;;
        --metrics)
            SHOW_METRICS=true
            SHOW_ALL=false
            shift
            ;;
        --streams)
            SHOW_STREAMS=true
            SHOW_ALL=false
            shift
            ;;
        --help|-h)
            echo "NATS Cluster Information Reporter"
            echo "Usage: $0 [options]"
            echo "Options:"
            echo "  --json     Output in JSON format"
            echo "  --metrics  Show only performance metrics"
            echo "  --streams  Show only stream information"
            echo "  --help     Show this help"
            exit 0
            ;;
        *)
            echo "Unknown option: $1"
            exit 1
            ;;
    esac
done

# Utility functions
format_bytes() {
    local bytes=$1
    if [[ $bytes -lt 1024 ]]; then
        echo "${bytes}B"
    elif [[ $bytes -lt 1048576 ]]; then
        echo "$(( bytes / 1024 ))KB"
    elif [[ $bytes -lt 1073741824 ]]; then
        echo "$(( bytes / 1048576 ))MB"
    else
        echo "$(( bytes / 1073741824 ))GB"
    fi
}

format_duration() {
    local seconds=$1
    if [[ $seconds -lt 60 ]]; then
        echo "${seconds}s"
    elif [[ $seconds -lt 3600 ]]; then
        echo "$(( seconds / 60 ))m $(( seconds % 60 ))s"
    elif [[ $seconds -lt 86400 ]]; then
        local hours=$(( seconds / 3600 ))
        local mins=$(( (seconds % 3600) / 60 ))
        echo "${hours}h ${mins}m"
    else
        local days=$(( seconds / 86400 ))
        local hours=$(( (seconds % 86400) / 3600 ))
        echo "${days}d ${hours}h"
    fi
}

# JSON output functions
json_start() {
    if [[ "$OUTPUT_JSON" == "true" ]]; then
        echo "{"
    fi
}

json_end() {
    if [[ "$OUTPUT_JSON" == "true" ]]; then
        echo "}"
    fi
}

json_section() {
    if [[ "$OUTPUT_JSON" == "true" ]]; then
        echo "  \"$1\": {"
    else
        echo -e "\n${BOLD}${CYAN}=== $1 ===${NC}"
    fi
}

json_section_end() {
    if [[ "$OUTPUT_JSON" == "true" ]]; then
        echo "  },"
    fi
}

# Data collection functions
get_cluster_topology() {
    local json_output=""
    
    if [[ "$OUTPUT_JSON" == "true" ]]; then
        json_output="    \"nodes\": ["
    else
        echo -e "${BOLD}Cluster Topology${NC}"
        printf "%-15s %-10s %-15s %-10s %-12s %-15s\n" "Node" "Status" "Role" "Uptime" "Connections" "Version"
        printf "%-15s %-10s %-15s %-10s %-12s %-15s\n" "----" "------" "----" "------" "-----------" "-------"
    fi
    
    local node_count=0
    for node in "${NATS_NODES[@]}"; do
        local status="OFFLINE"
        local role="Unknown"
        local uptime="N/A"
        local connections="0"
        local version="N/A"
        local cluster_name="N/A"
        local server_id="N/A"
        
        if curl -sf "http://$node/varz" >/dev/null 2>&1; then
            local varz
            varz=$(curl -s "http://$node/varz" 2>/dev/null || echo "{}")
            
            status="ONLINE"
            uptime=$(echo "$varz" | jq -r '.uptime // "0s"' 2>/dev/null | sed 's/[^0-9]*$//')
            connections=$(echo "$varz" | jq -r '.connections // 0' 2>/dev/null)
            version=$(echo "$varz" | jq -r '.version // "unknown"' 2>/dev/null)
            cluster_name=$(echo "$varz" | jq -r '.cluster.name // "N/A"' 2>/dev/null)
            server_id=$(echo "$varz" | jq -r '.server_id // "N/A"' 2>/dev/null)
            
            if echo "$varz" | jq -r '.cluster.leader' 2>/dev/null | grep -q "true"; then
                role="LEADER"
            else
                role="FOLLOWER"
            fi
            
            # Convert uptime to human readable if it's a number
            if [[ "$uptime" =~ ^[0-9]+$ ]]; then
                uptime=$(format_duration "$uptime")
            fi
        fi
        
        if [[ "$OUTPUT_JSON" == "true" ]]; then
            [[ $node_count -gt 0 ]] && json_output+=","
            json_output+="\n      {
        \"node\": \"$node\",
        \"status\": \"$status\",
        \"role\": \"$role\",
        \"uptime\": \"$uptime\",
        \"connections\": $connections,
        \"version\": \"$version\",
        \"cluster_name\": \"$cluster_name\",
        \"server_id\": \"$server_id\"
      }"
        else
            # Color coding for status
            local status_colored="$status"
            if [[ "$status" == "ONLINE" ]]; then
                status_colored="${GREEN}$status${NC}"
            else
                status_colored="${RED}$status${NC}"
            fi
            
            # Color coding for role
            local role_colored="$role"
            if [[ "$role" == "LEADER" ]]; then
                role_colored="${YELLOW}$role${NC}"
            elif [[ "$role" == "FOLLOWER" ]]; then
                role_colored="${BLUE}$role${NC}"
            fi
            
            printf "%-25s %-20s %-25s %-10s %-12s %-15s\n" \
                "$node" "$status_colored" "$role_colored" "$uptime" "$connections" "$version"
        fi
        
        node_count=$((node_count + 1))
    done
    
    if [[ "$OUTPUT_JSON" == "true" ]]; then
        echo "$json_output"
        echo "    ]"
    fi
}

get_jetstream_info() {
    if [[ "$OUTPUT_JSON" == "true" ]]; then
        echo "    \"jetstream\": {"
    else
        echo -e "\n${BOLD}JetStream Status${NC}"
    fi
    
    # Find JetStream leader
    local js_leader="None"
    local js_enabled_nodes=0
    local total_memory=0
    local total_storage=0
    local total_api_requests=0
    
    for node in "${NATS_NODES[@]}"; do
        if ! curl -sf "http://$node/jsz" >/dev/null 2>&1; then
            continue
        fi
        
        local jsz
        jsz=$(curl -s "http://$node/jsz" 2>/dev/null || echo "{}")
        
        if echo "$jsz" | jq -e '.config' >/dev/null 2>&1; then
            js_enabled_nodes=$((js_enabled_nodes + 1))
            
            # Check if this is the meta leader
            if echo "$jsz" | jq -r '.meta.leader' 2>/dev/null | grep -q "true"; then
                js_leader="$node"
            fi
            
            # Accumulate stats
            local memory
            local storage
            local api_reqs
            memory=$(echo "$jsz" | jq -r '.memory // 0' 2>/dev/null)
            storage=$(echo "$jsz" | jq -r '.store // 0' 2>/dev/null)
            api_reqs=$(echo "$jsz" | jq -r '.api.total // 0' 2>/dev/null)
            
            total_memory=$((total_memory + memory))
            total_storage=$((total_storage + storage))
            total_api_requests=$((total_api_requests + api_reqs))
        fi
    done
    
    if [[ "$OUTPUT_JSON" == "true" ]]; then
        echo "      \"leader\": \"$js_leader\","
        echo "      \"enabled_nodes\": $js_enabled_nodes,"
        echo "      \"total_memory_usage\": $total_memory,"
        echo "      \"total_storage_usage\": $total_storage,"
        echo "      \"total_api_requests\": $total_api_requests"
        echo "    },"
    else
        echo "JetStream Leader: $js_leader"
        echo "Enabled Nodes: $js_enabled_nodes"
        echo "Memory Usage: $(format_bytes $total_memory)"
        echo "Storage Usage: $(format_bytes $total_storage)"
        echo "API Requests: $total_api_requests"
    fi
}

get_stream_leadership() {
    # Find available server
    local server=""
    for node in "${NATS_NODES[@]}"; do
        if nats server check --server="$node" >/dev/null 2>&1; then
            server="$node"
            break
        fi
    done
    
    if [[ -z "$server" ]]; then
        if [[ "$OUTPUT_JSON" == "true" ]]; then
            echo "    \"streams\": [],"
        else
            echo -e "\n${BOLD}Stream Leadership${NC}"
            echo "No available NATS server to query streams"
        fi
        return
    fi
    
    # Get stream list
    local streams
    streams=$(nats stream list --server="$server" 2>/dev/null || echo "")
    
    if [[ "$OUTPUT_JSON" == "true" ]]; then
        echo "    \"streams\": ["
    else
        echo -e "\n${BOLD}Stream Leadership${NC}"
        if [[ -z "$streams" ]]; then
            echo "No streams configured"
            return
        fi
        printf "%-20s %-15s %-10s %-10s %-15s %-12s\n" "Stream" "Leader" "Replicas" "Messages" "Storage" "Consumers"
        printf "%-20s %-15s %-10s %-10s %-15s %-12s\n" "------" "------" "--------" "--------" "-------" "---------"
    fi
    
    local stream_count=0
    while IFS= read -r stream; do
        if [[ -z "$stream" ]]; then continue; fi
        
        local stream_info
        stream_info=$(nats stream info "$stream" --server="$server" --json 2>/dev/null || echo "{}")
        
        if [[ -z "$stream_info" || "$stream_info" == "{}" ]]; then
            continue
        fi
        
        local leader="N/A"
        local replicas="1"
        local messages="0"
        local storage="0B"
        local consumers="0"
        
        # Extract stream information
        leader=$(echo "$stream_info" | jq -r '.cluster.leader // "N/A"' 2>/dev/null)
        replicas=$(echo "$stream_info" | jq -r '.config.num_replicas // 1' 2>/dev/null)
        messages=$(echo "$stream_info" | jq -r '.state.messages // 0' 2>/dev/null)
        local bytes
        bytes=$(echo "$stream_info" | jq -r '.state.bytes // 0' 2>/dev/null)
        storage=$(format_bytes "$bytes")
        consumers=$(echo "$stream_info" | jq -r '.state.consumer_count // 0' 2>/dev/null)
        
        if [[ "$OUTPUT_JSON" == "true" ]]; then
            [[ $stream_count -gt 0 ]] && echo ","
            echo -n "      {
        \"name\": \"$stream\",
        \"leader\": \"$leader\",
        \"replicas\": $replicas,
        \"messages\": $messages,
        \"bytes\": $bytes,
        \"consumers\": $consumers
      }"
        else
            printf "%-20s %-15s %-10s %-10s %-15s %-12s\n" \
                "$stream" "$leader" "$replicas" "$messages" "$storage" "$consumers"
        fi
        
        stream_count=$((stream_count + 1))
    done <<< "$streams"
    
    if [[ "$OUTPUT_JSON" == "true" ]]; then
        [[ $stream_count -gt 0 ]] && echo
        echo "    ],"
    fi
}

get_performance_metrics() {
    if [[ "$OUTPUT_JSON" == "true" ]]; then
        echo "    \"performance\": {"
    else
        echo -e "\n${BOLD}Performance Metrics${NC}"
        printf "%-15s %-12s %-12s %-12s %-12s %-12s\n" "Node" "In Msgs/sec" "Out Msgs/sec" "In Bytes/sec" "Out Bytes/sec" "CPU %"
        printf "%-15s %-12s %-12s %-12s %-12s %-12s\n" "----" "-----------" "------------" "------------" "-------------" "-----"
    fi
    
    local total_in_msgs=0
    local total_out_msgs=0
    local total_in_bytes=0
    local total_out_bytes=0
    local node_count=0
    
    for node in "${NATS_NODES[@]}"; do
        if ! curl -sf "http://$node/varz" >/dev/null 2>&1; then
            if [[ "$OUTPUT_JSON" != "true" ]]; then
                printf "%-15s %-12s %-12s %-12s %-12s %-12s\n" "$node" "OFFLINE" "OFFLINE" "OFFLINE" "OFFLINE" "OFFLINE"
            fi
            continue
        fi
        
        local varz
        varz=$(curl -s "http://$node/varz" 2>/dev/null || echo "{}")
        
        local in_msgs
        local out_msgs
        local in_bytes
        local out_bytes
        local cpu_percent="N/A"
        
        in_msgs=$(echo "$varz" | jq -r '.in_msgs // 0' 2>/dev/null)
        out_msgs=$(echo "$varz" | jq -r '.out_msgs // 0' 2>/dev/null)
        in_bytes=$(echo "$varz" | jq -r '.in_bytes // 0' 2>/dev/null)
        out_bytes=$(echo "$varz" | jq -r '.out_bytes // 0' 2>/dev/null)
        cpu_percent=$(echo "$varz" | jq -r '.cpu // 0' 2>/dev/null)
        
        if [[ "$OUTPUT_JSON" == "true" ]]; then
            [[ $node_count -gt 0 ]] && echo ","
            echo -n "      \"$node\": {
        \"in_msgs_per_sec\": $in_msgs,
        \"out_msgs_per_sec\": $out_msgs,
        \"in_bytes_per_sec\": $in_bytes,
        \"out_bytes_per_sec\": $out_bytes,
        \"cpu_percent\": $cpu_percent
      }"
        else
            printf "%-15s %-12s %-12s %-12s %-12s %-12.1f\n" \
                "$node" "$in_msgs" "$out_msgs" "$(format_bytes $in_bytes)" "$(format_bytes $out_bytes)" "$cpu_percent"
        fi
        
        total_in_msgs=$((total_in_msgs + in_msgs))
        total_out_msgs=$((total_out_msgs + out_msgs))
        total_in_bytes=$((total_in_bytes + in_bytes))
        total_out_bytes=$((total_out_bytes + out_bytes))
        node_count=$((node_count + 1))
    done
    
    if [[ "$OUTPUT_JSON" == "true" ]]; then
        [[ $node_count -gt 0 ]] && echo ","
        echo "      \"totals\": {
        \"total_in_msgs_per_sec\": $total_in_msgs,
        \"total_out_msgs_per_sec\": $total_out_msgs,
        \"total_in_bytes_per_sec\": $total_in_bytes,
        \"total_out_bytes_per_sec\": $total_out_bytes
      }"
        echo "    },"
    else
        echo
        echo "Cluster Totals:"
        echo "  In Messages/sec: $total_in_msgs"
        echo "  Out Messages/sec: $total_out_msgs" 
        echo "  In Bytes/sec: $(format_bytes $total_in_bytes)"
        echo "  Out Bytes/sec: $(format_bytes $total_out_bytes)"
    fi
}

get_storage_usage() {
    if [[ "$OUTPUT_JSON" == "true" ]]; then
        echo "    \"storage\": ["
    else
        echo -e "\n${BOLD}Storage Usage${NC}"
        printf "%-15s %-15s %-15s %-15s\n" "Container" "Total Size" "Used Space" "Available"
        printf "%-15s %-15s %-15s %-15s\n" "---------" "----------" "----------" "---------"
    fi
    
    local container_count=0
    for container in "${NATS_CONTAINERS[@]}"; do
        if ! docker ps --format '{{.Names}}' | grep -q "^${container}$"; then
            if [[ "$OUTPUT_JSON" != "true" ]]; then
                printf "%-15s %-15s %-15s %-15s\n" "$container" "OFFLINE" "OFFLINE" "OFFLINE"
            fi
            continue
        fi
        
        # Get filesystem usage for JetStream data directory
        local df_output
        df_output=$(docker exec "$container" df -h /data 2>/dev/null || echo "")
        
        if [[ -n "$df_output" ]]; then
            local size used avail
            size=$(echo "$df_output" | tail -n1 | awk '{print $2}')
            used=$(echo "$df_output" | tail -n1 | awk '{print $3}')
            avail=$(echo "$df_output" | tail -n1 | awk '{print $4}')
            
            if [[ "$OUTPUT_JSON" == "true" ]]; then
                [[ $container_count -gt 0 ]] && echo ","
                echo -n "      {
        \"container\": \"$container\",
        \"total_size\": \"$size\",
        \"used_space\": \"$used\",
        \"available\": \"$avail\"
      }"
            else
                printf "%-15s %-15s %-15s %-15s\n" "$container" "$size" "$used" "$avail"
            fi
        else
            if [[ "$OUTPUT_JSON" != "true" ]]; then
                printf "%-15s %-15s %-15s %-15s\n" "$container" "N/A" "N/A" "N/A"
            fi
        fi
        
        container_count=$((container_count + 1))
    done
    
    if [[ "$OUTPUT_JSON" == "true" ]]; then
        [[ $container_count -gt 0 ]] && echo
        echo "    ]"
    fi
}

# Main execution
main() {
    if [[ "$OUTPUT_JSON" == "true" ]]; then
        json_start
        echo "  \"timestamp\": \"$(date -u +"%Y-%m-%dT%H:%M:%SZ")\","
        echo "  \"cluster_name\": \"nats-cluster\","
    else
        echo -e "${BOLD}${CYAN}NATS Cluster Information Report${NC}"
        echo "Generated: $(date)"
        echo "Cluster: nats-cluster"
    fi
    
    if [[ "$SHOW_ALL" == "true" || "$SHOW_STREAMS" == "false" ]]; then
        get_cluster_topology
        get_jetstream_info
        get_stream_leadership
    fi
    
    if [[ "$SHOW_ALL" == "true" || "$SHOW_METRICS" == "true" ]]; then
        get_performance_metrics
        get_storage_usage
    fi
    
    if [[ "$SHOW_STREAMS" == "true" && "$SHOW_ALL" == "false" ]]; then
        get_stream_leadership
    fi
    
    if [[ "$OUTPUT_JSON" == "true" ]]; then
        # Remove trailing comma from last section
        echo "  \"report_complete\": true"
        json_end
    else
        echo -e "\n${GREEN}Report complete.${NC}"
    fi
}

# Execute main function
main "$@"