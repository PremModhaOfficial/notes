# Hour 2 Progress Tracker  
## Docker Setup + 3-Node Cluster

> **MODULE**: Hour 2 of 5  
> **DURATION**: 60 minutes  
> **STATUS**: 🔄 Ready to Begin  
> **LAST UPDATED**: 2026-01-07  

---

## 📚 Learning Materials

- [[NATS_H2_DOCKER]] - Docker Setup Guide (30 min)
- [[NATS_H2_CLUSTER]] - Cluster Configuration Deep Dive (30 min)

---

## 🎯 Learning Objectives

By the end of Hour 2, you will have:

### Infrastructure Checkboxes:
- [ ] **Single NATS node** running with JetStream enabled
- [ ] **3-node cluster** formed and healthy  
- [ ] **Docker networking** understood for inter-node communication
- [ ] **Configuration files** created and explained
- [ ] **Verification commands** working to check cluster state
- [ ] **Basic troubleshooting** skills for common issues

### Technical Understanding:
- [ ] **Know the difference** between client ports (4222) and cluster ports (6222)
- [ ] **Understand routes** - how nodes discover each other
- [ ] **Monitor cluster health** via HTTP endpoints (/varz, /routez)
- [ ] **JetStream persistence** with Docker volumes
- [ ] **Container networking** concepts for clustering

---

## 🛠️ Prerequisites

Before starting Hour 2:
- [ ] **Hour 1 completed** - RAFT theory understood
- [ ] **Docker installed** - `docker --version` works
- [ ] **Docker Compose installed** - `docker compose version` works  
- [ ] **Terminal ready** - comfortable with command line
- [ ] **60 minutes available** - focused work session

---

## ⏱️ Time Breakdown

| Minutes | Activity | Checkpoint |
|---------|----------|------------|
| 0-10 | Environment setup + single node test | 🟢 Docker working |
| 10-30 | Read [[NATS_H2_DOCKER]] and follow along | 🟡 Single node confirmed |
| 30-50 | Build 3-node cluster + verification | 🟡 Cluster formed |
| 50-60 | Read [[NATS_H2_CLUSTER]] + troubleshoot | 🔴 Production ready |

---

## 🔧 File Structure Created

By the end of this hour, you'll have:

```
configs/
├── NATS_docker-compose.yml          # 3-node cluster
├── NATS_docker-compose.single.yml   # Single node for testing  
├── NATS_nats-1.conf                 # First node config
├── NATS_nats-2.conf                 # Second node config
└── NATS_nats-3.conf                 # Third node config
```

---

## 🏥 Health Check Commands

Learn these verification commands:

```bash
# Check single node
curl http://localhost:8222/healthz

# Check cluster formation  
curl http://localhost:8222/varz | jq '.cluster'

# Check node routes
curl http://localhost:8222/routez | jq '.routes[].remote_name'

# Check JetStream status
curl http://localhost:8222/jsz | jq '.'
```

---

## 🐛 Common Issues & Solutions

| Problem | Solution |
|---------|----------|
| "Connection refused" | Check if container is running: `docker ps` |
| "Cluster not forming" | Verify hostnames match in routes config |
| "JetStream disabled" | Check `jetstream {}` block in config |
| Port conflicts | Use different host ports (4222, 4223, 4224) |

---

## 🚦 Session Status

### Current State:
- **Files Created**: Docker configs ready (via background agent)
- **Cluster Status**: Not started
- **Dependencies**: Hour 1 theory completed
- **Next Step**: Begin Docker setup

### Success Criteria:
- [ ] All 3 NATS containers running  
- [ ] Cluster health check passes
- [ ] Can connect to any node on different ports
- [ ] JetStream enabled on all nodes
- [ ] Ready for Hour 3 JetStream work

---

## 🔄 For AI Assistants

### Handoff Protocol:
1. **Check Prerequisites**: Verify Docker installation
2. **Follow Materials**: Use [[NATS_H2_DOCKER]] as primary guide  
3. **Troubleshoot**: Reference common issues above
4. **Update Progress**: Mark checkboxes as completed
5. **Prepare Hour 3**: Ensure cluster is stable

### Debug Commands:
```bash
# If stuck, run these diagnostics
docker compose logs nats-1
docker compose ps
docker network ls | grep nats
```

---

*Previous: [[NATS_H1_PROGRESS]] | Next: [[NATS_H3_PROGRESS]] →*