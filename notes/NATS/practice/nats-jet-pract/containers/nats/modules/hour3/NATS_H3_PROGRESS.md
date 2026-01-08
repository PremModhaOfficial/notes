# Hour 3 Progress Tracker
## JetStream + Go Client Integration

> **MODULE**: Hour 3 of 5  
> **DURATION**: 60 minutes  
> **STATUS**: 🔄 Ready to Begin  
> **LAST UPDATED**: 2026-01-07  

---

## 📚 Learning Materials

- [[NATS_H3_JETSTREAM]] - JetStream Streams & RAFT Groups (30 min)
- [[NATS_H3_GO_CLIENT]] - Go Client Application (30 min)

---

## 🎯 Learning Objectives  

By the end of Hour 3, you will:

### JetStream Mastery:
- [ ] **Create replicated streams** with `replicas=3` setting
- [ ] **Understand multiple RAFT groups** - Meta vs Stream groups
- [ ] **See leader election per stream** - different leaders for different streams
- [ ] **Create durable consumers** that survive node failures
- [ ] **Monitor replica health** using NATS CLI and HTTP endpoints
- [ ] **Understand stream placement** across cluster nodes

### Go Client Skills:
- [ ] **Connect to NATS cluster** with automatic node discovery
- [ ] **Handle reconnections** gracefully during node failures
- [ ] **Publish to JetStream** with acknowledgments
- [ ] **Consume messages** with proper acknowledgment patterns
- [ ] **See client behavior** during leader elections
- [ ] **Graceful shutdown** handling

---

## 🛠️ Prerequisites

Before starting Hour 3:
- [ ] **Hour 2 completed** - 3-node cluster running and healthy
- [ ] **NATS CLI installed** - `go install github.com/nats-io/natscli/nats@latest`
- [ ] **Go environment** - `go version` shows 1.21+
- [ ] **Cluster verification** - `curl http://localhost:8222/varz` works
- [ ] **All 3 nodes accessible** - ports 4222, 4223, 4224 responding

---

## ⏱️ Time Breakdown

| Minutes | Activity | Checkpoint |
|---------|----------|------------|
| 0-15 | Install NATS CLI + create first stream | 🟢 Tools ready |
| 15-30 | Read [[NATS_H3_JETSTREAM]] + hands-on practice | 🟡 Streams replicated |
| 30-45 | Set up Go client + run publisher/consumer | 🟡 Client connected |
| 45-60 | Read [[NATS_H3_GO_CLIENT]] + test failover | 🔴 Resilience proven |

---

## 🏗️ What You'll Build

By the end of this hour:

### JetStream Resources:
```bash
# Streams created
ORDERS (replicas=3, leader on nats-1)
LOGS (replicas=3, leader on nats-2)  
EVENTS (replicas=3, leader on nats-3)

# Consumers created  
order-processor (durable, on ORDERS stream)
log-analyzer (durable, on LOGS stream)
```

### Go Application:
- **Publisher**: Sends orders every 2 seconds to multiple streams
- **Consumer**: Processes messages with proper acknowledgment
- **Monitoring**: Shows which NATS node it's connected to
- **Resilience**: Handles node failures gracefully

---

## 🔍 Key Concepts to Observe

### Multiple RAFT Groups in Action:
- **Meta RAFT Group**: Tracks which streams exist (all 3 nodes participate)
- **ORDERS RAFT Group**: Replicates order messages (3 nodes, own leader)
- **LOGS RAFT Group**: Replicates log messages (3 nodes, different leader)

### Client Behavior Patterns:
- **Discovery**: Client finds all cluster nodes automatically
- **Load Balancing**: Publishes are distributed across nodes  
- **Failover**: Client reconnects to healthy nodes during failures
- **Acknowledgments**: JetStream ensures messages are replicated before ACK

---

## 🧪 Experiments to Try

### Stream Leadership Experiment:
```bash
# Check which node leads each stream
nats stream info ORDERS | grep Leader
nats stream info LOGS | grep Leader  

# Kill the ORDERS leader, watch election
docker stop nats-1  # (assuming it was ORDERS leader)
nats stream info ORDERS  # New leader elected!
```

### Client Reconnection Experiment:
```bash
# Start Go client, watch output
go run go-client/NATS_main.go

# Kill node it's connected to
docker stop nats-2

# Watch client reconnect to nats-1 or nats-3
# Messages continue flowing!
```

---

## 🚦 Session Status

### Success Criteria:
- [ ] **3 replicated streams** created and healthy
- [ ] **Go client running** and publishing/consuming
- [ ] **Failover tested** - killed leader, watched election  
- [ ] **Client resilience** - survived node failures
- [ ] **RAFT groups understood** - multiple leaders for different streams
- [ ] **Ready for chaos testing** in Hour 4

### Files Created:
- [ ] `go-client/NATS_main.go` - Working Go application
- [ ] `go-client/NATS_go.mod` - Dependencies configured
- [ ] `go-client/NATS_README.md` - Usage documentation

---

## 🔄 For AI Assistants

### Troubleshooting Hints:
```bash
# If NATS CLI not working
export PATH=$PATH:$(go env GOPATH)/bin

# If Go client can't connect
nats server list  # Verify cluster accessible

# If streams not replicating  
nats stream info STREAMNAME --json | jq '.cluster'
```

### Next Steps:
- Ensure all experiments completed successfully
- Verify client handles failures gracefully
- Prepare for Hour 4 chaos engineering
- Update progress checkboxes

---

*Previous: [[NATS_H2_PROGRESS]] | Next: [[NATS_H4_PROGRESS]] →*