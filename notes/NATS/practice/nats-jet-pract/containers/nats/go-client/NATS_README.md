# NATS RAFT Learning - Go Client

This Go client application demonstrates NATS JetStream functionality with cluster support, designed for the NATS RAFT learning program. It showcases connection handling, stream management, message publishing, consuming, and graceful shutdown patterns.

## Features

- **Cluster Discovery**: Automatically connects to available NATS nodes in a cluster
- **Reconnection Handling**: Robust reconnection logic with jitter and unlimited retries
- **JetStream Integration**: Creates streams and consumers programmatically
- **Message Publishing**: Publishes structured messages every 2 seconds
- **Message Consuming**: Processes messages with proper acknowledgment
- **Connection Monitoring**: Displays connection status and statistics
- **Graceful Shutdown**: Clean shutdown handling with proper resource cleanup

## Prerequisites

- Go 1.21 or higher
- Docker and Docker Compose (for running NATS cluster)
- NATS cluster running (single node or multi-node)

## Quick Start

### 1. Start NATS Cluster

Choose one of the following configurations:

**Option A: Single Node (Development)**
```bash
# From the project root directory
docker-compose -f configs/NATS_docker-compose.single.yml up -d
```

**Option B: 3-Node Cluster (Production-like)**
```bash
# From the project root directory  
docker-compose -f configs/NATS_docker-compose.yml up -d
```

### 2. Run the Go Client

```bash
# Navigate to the go-client directory
cd go-client

# Download dependencies
go mod tidy

# Run the client
go run NATS_main.go
```

## What It Demonstrates

### Connection Management
- **Multi-server URLs**: Attempts connection to multiple NATS nodes
- **Automatic Discovery**: Discovers and connects to available cluster members
- **Reconnection Logic**: Handles disconnections with exponential backoff and jitter
- **Connection Events**: Logs disconnect, reconnect, and close events

### JetStream Features
- **Stream Creation**: Creates a persistent stream named `LEARNING_STREAM`
- **Consumer Setup**: Configures a durable consumer `LEARNING_CONSUMER`
- **Message Persistence**: Uses file-based storage with retention policies
- **Acknowledgment**: Demonstrates explicit message acknowledgment

### Message Patterns
- **Publishing**: Sends JSON messages with metadata every 2 seconds
- **Consuming**: Processes messages with proper error handling
- **Flow Control**: Manages message delivery and acknowledgment timing

## Expected Output

### Startup
```
🚀 Starting NATS RAFT Learning Client
=====================================
🔍 Attempting to connect to NATS cluster: [nats://localhost:4222 nats://localhost:4223 nats://localhost:4224]
✅ Connected to NATS server: nats://127.0.0.1:4222
📋 Server ID: NDJE5QZJPVDV7ZXHGYUFEWKKBZ2WPCYB7JHDQHZXT7VH7QIVPBOVKC4W
🏷️  Server Name: nats-1
🎯 JetStream context created successfully
🏗️  Initializing JetStream resources...
✅ Stream 'LEARNING_STREAM' created/updated successfully
📊 Stream info - Messages: 0, Bytes: 0, Consumers: 0
✅ Consumer 'LEARNING_CONSUMER' created/updated successfully
🎬 Starting client operations...
📤 Starting message publisher...
📥 Starting message consumer...
✅ Client started - publisher and consumer are running
📥 Consumer ready - waiting for messages...
```

### Runtime Messages
```
📤 Published message 1 (seq: 1, stream: LEARNING_STREAM) to learning.messages
📥 Received message (seq: 1, delivered: 1 times)
   Subject: learning.messages
   Data: {
		"id": 1,
		"timestamp": "2026-01-07T12:30:15Z",
		"message": "Hello from NATS RAFT Learning Client",
		"node": "nats://127.0.0.1:4222",
		"counter": 1
	}
   Timestamp: 2026-01-07T12:30:15Z
✅ Message acknowledged successfully
---
```

### Reconnection (when nodes fail)
```
⚠️  Disconnected from NATS: nats: connection closed
🔄 Reconnected to NATS server: nats://127.0.0.1:4223
📊 Connection stats - Reconnects: 1
```

### Graceful Shutdown
```
🛑 Shutdown signal received, cleaning up...
🛑 Initiating graceful shutdown...
📤 Publisher shutting down...
📥 Consumer shutting down...
✅ All goroutines stopped gracefully
🔐 Closing NATS connection...
✅ NATS connection closed
✅ Client shutdown complete
```

## Configuration

### Stream Configuration
- **Name**: `LEARNING_STREAM`
- **Subjects**: `learning.*`
- **Retention**: WorkQueue policy (messages removed after ACK)
- **Storage**: File-based persistent storage
- **Max Messages**: 1,000
- **Max Age**: 24 hours

### Consumer Configuration
- **Name**: `LEARNING_CONSUMER`
- **Type**: Durable (survives restarts)
- **Ack Policy**: Explicit acknowledgment required
- **Max Deliver**: 3 attempts
- **Ack Wait**: 30 seconds before redelivery

### Connection Settings
- **Reconnect**: Unlimited attempts
- **Reconnect Wait**: 2 seconds with jitter
- **Timeout**: 10 seconds for operations

## Troubleshooting

### Common Issues

**1. Connection Failed**
```
❌ Failed to connect to any NATS server: dial tcp [::1]:4222: connect: connection refused
```
**Solution**: Ensure NATS cluster is running:
```bash
docker ps | grep nats
docker-compose -f configs/NATS_docker-compose.single.yml up -d
```

**2. JetStream Not Available**
```
❌ Failed to create JetStream context: nats: JetStream not enabled
```
**Solution**: Verify JetStream is enabled in NATS configuration (should be enabled in provided configs)

**3. Permission Errors**
```
❌ Failed to create stream: nats: insufficient permissions
```
**Solution**: This example uses default NATS (no auth). Check if your NATS server requires authentication.

**4. Go Module Issues**
```
go: module github.com/nats-io/nats.go: Get "https://proxy.golang.org/github.com/nats-io/nats.go/@v/list": dial tcp: i/o timeout
```
**Solution**: 
```bash
go clean -modcache
go mod tidy
```

### Testing Node Failures

**Single Node Environment**:
```bash
# Stop single node
docker stop nats-single

# Restart node  
docker start nats-single
```

**Cluster Environment**:
```bash
# Stop one node to test failover
docker stop nats-1

# Stop majority to test split-brain handling
docker stop nats-1 nats-2

# Restart nodes
docker start nats-1 nats-2
```

### Monitoring

**NATS Monitoring UI**:
- Single Node: http://localhost:8222
- Node 1: http://localhost:8222  
- Node 2: http://localhost:8223
- Node 3: http://localhost:8224

**View JetStream Info**:
```bash
# Install NATS CLI
go install github.com/nats-io/natscli/nats@latest

# View stream info
nats stream info LEARNING_STREAM

# View consumer info
nats consumer info LEARNING_STREAM LEARNING_CONSUMER
```

## Learning Objectives

This client demonstrates:

1. **RAFT Consensus**: How JetStream uses RAFT for leader election and log replication
2. **High Availability**: Client behavior during node failures and recovery
3. **Message Persistence**: How messages survive node restarts with JetStream
4. **Load Balancing**: How clients distribute across cluster nodes
5. **Split-Brain Handling**: Behavior when majority of nodes are unavailable

## Code Structure

- `NATSClient`: Main client structure with all components
- `NewNATSClient()`: Connection setup with cluster discovery
- `initializeJetStream()`: Stream and consumer creation
- `runPublisher()`: Message publishing loop
- `runConsumer()`: Message consumption and processing
- `Shutdown()`: Graceful shutdown handling

## Usage in NATS RAFT Learning

This client is designed for **Hour 3** of the NATS RAFT learning program to demonstrate:
- JetStream replication across cluster nodes
- Client behavior during controlled node failures
- Message persistence and recovery scenarios
- Cluster formation and split-brain resolution

Run the client while performing cluster operations to observe real-time behavior and understand NATS RAFT consensus in action.