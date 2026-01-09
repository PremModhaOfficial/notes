# NATS Micro Service Framework - Deep Dive

**Time**: ~2 hours (Part 1: 1 hour, Part 2: 1 hour)  
**Prerequisites**: Basic NATS pub/sub, request/reply pattern, Go fundamentals  
**Related**: [[../nats-consumer/README|NATS Consumers Module]]

---

## Quick Start

```bash
docker compose up -d
cd part-1-foundations/examples/01-basic-service
go run main.go
```

In another terminal:
```bash
docker exec -it nats-micro-lab nats req echo "hello world"
docker exec -it nats-micro-lab nats micro ls
```

---

## Learning Path

### Part 1: Foundations (Hour 3)
**[→ Start Part 1](./part-1-foundations/README.md)**

- What is the Micro Framework?
- Service Configuration
- Handlers & Requests
- Endpoints & Groups
- Queue Groups
- Discovery Protocol ($SRV)

### Part 2: Internals & Patterns (Hour 4)
**[→ Start Part 2](./part-2-internals-patterns/README.md)**

- $SRV Protocol Internals
- Service Lifecycle
- Stats Collection
- Error Handling
- Production Patterns
- Micro + JetStream Integration

---

## What You'll Learn

After completing this module, you will:

1. **Build** microservices using the NATS micro framework
2. **Understand** the $SRV discovery protocol internals
3. **Implement** proper error handling and graceful shutdown
4. **Monitor** services with stats and custom metrics
5. **Design** hybrid architectures combining Micro and JetStream

---

## File Structure

```
nats-micro/
├── README.md                           # You are here
├── docker-compose.yml                  # Lab environment
├── part-1-foundations/
│   ├── README.md                       # Main learning doc
│   ├── diagrams/                       # ASCII diagrams
│   ├── examples/                       # Runnable Go code
│   └── exercises.md                    # Hands-on challenges
└── part-2-internals-patterns/
    ├── README.md                       # Deep internals doc
    ├── diagrams/
    ├── examples/
    └── exercises.md
```

---

## Cross-References

- **Consumers**: See [[../nats-consumer/README|NATS Consumers Module]] for JetStream consumption
- **RAFT Consensus**: See [[../nats/modules/hour1/NATS_H1_THEORY|Hour 1 RAFT Theory]] for cluster consensus
- **Cluster Setup**: See [[../nats/modules/hour2/NATS_H2_DOCKER|Hour 2 Docker]] for multi-node testing
