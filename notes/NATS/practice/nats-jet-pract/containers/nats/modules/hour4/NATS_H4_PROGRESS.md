# Hour 4 Progress Tracker
## Chaos Testing + Production Tuning

> **MODULE**: Hour 4 of 5  
> **DURATION**: 60 minutes  
> **STATUS**: 🔄 Ready to Begin  
> **LAST UPDATED**: 2026-01-07  

---

## 📚 Learning Materials

- [[NATS_H4_CHAOS]] - Systematic Failure Testing (40 min)
- [[NATS_H4_PRODUCTION]] - Production Best Practices (20 min)

---

## 🎯 Learning Objectives

By the end of Hour 4, you will:

### Chaos Engineering Mastery:
- [ ] **Execute systematic failure scenarios** using automated scripts
- [ ] **Measure RAFT election timing** (should be 2-3 seconds)
- [ ] **Observe split-brain prevention** when majority is lost
- [ ] **Test network partition recovery** when connectivity returns
- [ ] **Verify zero message loss** during leader failures
- [ ] **Understand recovery patterns** and cluster healing

### Production Readiness:
- [ ] **Optimize configuration** for production workloads
- [ ] **Implement security hardening** basics
- [ ] **Set up backup procedures** for JetStream data
- [ ] **Configure monitoring** for operational awareness
- [ ] **Document runbook procedures** for operations team

---

## 🛠️ Prerequisites

Before starting Hour 4:
- [ ] **Hour 3 completed** - Go client working with replicated streams
- [ ] **Cluster stable** - all 3 nodes healthy
- [ ] **Message flow confirmed** - publisher/consumer working
- [ ] **NATS CLI functional** - can create/inspect streams
- [ ] **Basic shell scripting** - comfortable running bash scripts

---

## ⏱️ Time Breakdown

| Minutes | Activity | Checkpoint |
|---------|----------|------------|
| 0-10 | Run baseline health checks + setup | 🟢 Baseline established |
| 10-30 | Execute chaos scenarios 1-3 | 🟡 Failure patterns understood |
| 30-40 | Measure and document recovery timing | 🟡 Performance characterized |
| 40-55 | Production tuning + security basics | 🔴 Production ready |
| 55-60 | Backup procedures + runbook creation | 🔴 Operational handoff complete |

---

## 🧪 Chaos Testing Scenarios

### Scenario 1: Leader Election 🗳️
**Test**: Kill current stream leader, watch election
- **Expected**: New leader elected in 2-3 seconds
- **Verify**: Zero message loss, client reconnects automatically
- **Tools**: `NATS_chaos-test.sh --scenario=leader-failure`

### Scenario 2: Split Brain Prevention 🧠
**Test**: Kill 2 out of 3 nodes (lose quorum)
- **Expected**: Cluster becomes read-only, no new writes accepted
- **Verify**: Existing data preserved, cluster waits for healing
- **Tools**: `NATS_chaos-test.sh --scenario=split-brain`

### Scenario 3: Network Partition Recovery 🔗
**Test**: Simulate network partition, then heal
- **Expected**: Partitioned nodes catch up when reconnected
- **Verify**: Log reconciliation, eventual consistency
- **Tools**: Docker network manipulation

### Scenario 4: Cascading Failures ⛓️
**Test**: Kill nodes in sequence, test recovery under load
- **Expected**: Graceful degradation, no data corruption
- **Verify**: Client handles multiple reconnections

---

## 📊 Key Metrics to Track

### Performance Metrics:
- **Leader Election Time**: Should be < 3 seconds
- **Message Throughput**: Baseline vs during failures
- **Client Reconnection Time**: Should be sub-second
- **Log Replication Lag**: Monitor during catch-up

### Health Indicators:
- **Cluster Quorum Status**: 2/3 or 3/3 healthy
- **Stream Replica Health**: All replicas current
- **Consumer Lag**: Messages pending processing
- **Storage Usage**: Disk space per node

---

## 🏭 Production Configuration Changes

### Key Optimizations:
- [ ] **Memory allocation**: Increase `max_memory_store` to 2-8GB
- [ ] **File storage**: Set `max_file_store` to 50-500GB
- [ ] **Connection limits**: Tune `max_connections` for workload
- [ ] **Heartbeat timing**: Optimize RAFT election timeouts
- [ ] **Log retention**: Configure appropriate message retention

### Security Hardening:
- [ ] **Authentication**: Enable user/password or JWT
- [ ] **TLS encryption**: Secure client and cluster connections
- [ ] **Authorization**: Implement subject-level permissions
- [ ] **Network isolation**: Restrict cluster ports access

---

## 🔧 Tools Created

### Automation Scripts:
- [ ] `scripts/NATS_chaos-test.sh` - Automated failure testing
- [ ] `scripts/NATS_health-check.sh` - Continuous health monitoring  
- [ ] `scripts/NATS_cluster-info.sh` - Detailed cluster reporting

### Usage Examples:
```bash
# Run full chaos test suite
./scripts/NATS_chaos-test.sh --all

# Quick health check
./scripts/NATS_health-check.sh

# Detailed cluster report
./scripts/NATS_cluster-info.sh --verbose
```

---

## 🚦 Session Status

### Success Criteria:
- [ ] **All chaos scenarios passed** - cluster survived systematic failures
- [ ] **Recovery timing documented** - know your RTO/RPO metrics
- [ ] **Production config optimized** - ready for real workloads
- [ ] **Security basics implemented** - authentication and encryption
- [ ] **Operational procedures documented** - team can maintain system

### Failure Modes Tested:
- [ ] **Leader failure** → New election in < 3 seconds
- [ ] **Majority loss** → Read-only mode, data preserved  
- [ ] **Network partition** → Automatic healing when connectivity restored
- [ ] **Cascading failures** → Graceful degradation maintained

---

## 📚 Documentation Outputs

By end of Hour 4:
- [ ] **Chaos test results** - timing and behavior documentation
- [ ] **Performance baseline** - throughput and latency metrics
- [ ] **Production checklist** - configuration and security items
- [ ] **Runbook draft** - operational procedures for team

---

## 🔄 For AI Assistants

### Testing Protocol:
1. **Establish baseline** using health check scripts
2. **Execute scenarios** systematically, document results
3. **Measure timing** for all recovery operations
4. **Verify data integrity** after each test
5. **Update configurations** based on findings

### Common Issues:
- **Split brain confusion**: Remember, this is GOOD behavior
- **Slow elections**: May indicate network or resource issues
- **Client errors**: Expected during failures, should recover automatically

---

*Previous: [[NATS_H3_PROGRESS]] | Next: [[NATS_H5_PROGRESS]] →*