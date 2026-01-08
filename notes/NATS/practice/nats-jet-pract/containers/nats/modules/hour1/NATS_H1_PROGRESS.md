# Hour 1 Progress Tracker
## RAFT Theory + NATS Concepts

> **MODULE**: Hour 1 of 5  
> **DURATION**: 60 minutes  
> **STATUS**: 🔄 Ready to Begin  
> **LAST UPDATED**: 2026-01-07  

---

## 📚 Learning Materials

- [[NATS_H1_THEORY]] - RAFT Consensus Algorithm (30 min)
- [[NATS_H1_CONCEPTS]] - NATS Architecture & JetStream (30 min)

---

## 🎯 Learning Objectives

By the end of Hour 1, you will:

### RAFT Mastery Checkboxes:
- [ ] **Explain the consensus problem** using the pizza shop analogy
- [ ] **Visualize the 3 RAFT states**: Follower → Candidate → Leader  
- [ ] **Understand leader election** with randomized timeouts
- [ ] **Know quorum mathematics**: 3 nodes = 2 quorum, 5 nodes = 3 quorum
- [ ] **Trace log replication**: Leader → Followers → Commit → Client ACK
- [ ] **Predict failure scenarios**: What happens when leader dies?

### NATS Architecture Checkboxes:
- [ ] **Distinguish NATS Core vs JetStream** - messaging vs persistence
- [ ] **Understand multiple RAFT groups** - Meta, Stream, Consumer groups
- [ ] **Appreciate NATS RAFT optimization** - data plane + control plane merged
- [ ] **Know cluster topology options** - single cluster vs supercluster  
- [ ] **Understand client behavior** - discovery, reconnection, load balancing

---

## ⏱️ Time Breakdown

| Minutes | Activity | Checkpoint |
|---------|----------|------------|
| 0-30 | Read [[NATS_H1_THEORY]] | 🟢 RAFT concepts clear |
| 30-45 | Read [[NATS_H1_CONCEPTS]] | 🟡 NATS architecture understood |
| 45-60 | Interactive exercises + prep for Hour 2 | 🔴 Ready for hands-on work |

---

## 🧠 Knowledge Check

Before moving to Hour 2, ensure you can answer:

### Quick Quiz:
1. **Why do we need odd numbers of nodes?** (3, 5, 7 not 2, 4, 6)
2. **What happens if 2 out of 3 nodes die?** 
3. **How does NATS JetStream use multiple RAFT groups?**
4. **Why is NATS RAFT more efficient than traditional implementations?**

### Expected Answers:
1. Even numbers don't improve fault tolerance but add overhead
2. Cluster becomes read-only (no quorum for writes)  
3. Separate RAFT groups for meta-data, each stream, and consumers
4. Combines data replication with RAFT consensus messages

---

## 🚦 Session Status

### Prerequisites Met:
- [ ] Comfortable with basic distributed systems terminology
- [ ] Ready to dive deep into technical concepts
- [ ] Have 60 minutes for focused learning

### Post-Hour 1 Checklist:
- [ ] All learning objectives checkboxes completed
- [ ] Knowledge check questions answered confidently  
- [ ] Excited about building the actual cluster in Hour 2
- [ ] Environment verified for Docker work

---

## 🔄 For AI Assistants

### Current State: 
- **Module**: Hour 1 - Theory & Concepts
- **Status**: Ready to begin or continue from checkpoint
- **Dependencies**: None (pure learning)
- **Next**: Hour 2 Docker implementation

### Handoff Instructions:
1. If learner is stuck on concepts, revisit analogies in theory file
2. If ready to proceed, guide them to Hour 2 Docker setup
3. If knowledge check fails, review weak areas before continuing
4. Update status in this file after session

---

*Previous: [[NATS_RAFT_ROADMAP]] | Next: [[NATS_H2_PROGRESS]] →*