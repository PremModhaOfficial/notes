# NATS JetStream Consumers - Deep Dive

**Time**: ~2 hours (Part 1: 1 hour, Part 2: 1 hour)  
**Prerequisites**: Basic NATS pub/sub, JetStream streams, Go fundamentals  
**Related**: [[../nats/modules/hour3/NATS_H3_JETSTREAM|RAFT & JetStream Module]]

---

## Quick Start

```bash
docker compose up -d
docker exec -it nats-consumer-lab nats stream add ORDERS --subjects "orders.>" --defaults
docker exec -it nats-consumer-lab nats pub orders.new "test message"
```

---

## Learning Path

### Part 1: Foundations (Hour 1)
**[→ Start Part 1](./part-1-foundations/README.md)**

- What is a Consumer?
- Pull vs Push Consumers
- Durable vs Ephemeral
- Delivery Policies
- ACK Policies
- Filter Subjects

### Part 2: Internals & Debugging (Hour 2)
**[→ Start Part 2](./part-2-internals-debugging/README.md)**

- ACK Protocol Wire Format
- The 5 ACK Types
- Consumer State Machine
- AckWait & BackOff
- MaxAckPending Flow Control
- Debugging Common Issues

---

## What You'll Learn

After completing this module, you will:

1. **Understand** how consumers track message delivery and acknowledgment
2. **Choose** the right consumer type (pull/push, durable/ephemeral) for your use case
3. **Debug** consumer issues like stalls, redelivery storms, and flow control
4. **Tune** consumer configuration for production workloads
5. **Think** like a senior engineer when troubleshooting consumer problems

---

## File Structure

```
nats-consumer/
├── README.md                           # You are here
├── docker-compose.yml                  # Lab environment
├── part-1-foundations/
│   ├── README.md                       # Main learning doc
│   ├── diagrams/                       # ASCII diagrams
│   ├── examples/                       # Runnable Go code
│   └── exercises.md                    # Hands-on challenges
└── part-2-internals-debugging/
    ├── README.md                       # Deep internals doc
    ├── diagrams/
    ├── examples/
    └── exercises.md
```

---

## Cross-References

- **RAFT Consensus**: See [[../nats/modules/hour1/NATS_H1_THEORY|Hour 1 RAFT Theory]] for how consumer state is replicated
- **Cluster Setup**: See [[../nats/modules/hour2/NATS_H2_DOCKER|Hour 2 Docker]] for multi-node testing
- **Micro Services**: See [[../nats-micro/README|NATS Micro Module]] for request/reply patterns
