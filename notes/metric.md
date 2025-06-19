---
id: metric
aliases:
  - Metric
tags:
  - topic/Observability
  - concept/devops
  - source/udemy
created: 2025-06-17 11:30:54
modified: 2025-06-19 10:32:50
status:
  - #draft
type:
  - #atomic-note
---

# Metric
## 🎯 Core Concept
### The Problem
- the data in our system is very large
  * The data of thread state transitions
  * Connections
  * The names of the open files etc
  * CPU states

### The solution
- we only store the aggrigates
  * Connections+ Threds: the number of running, waiting and Max numbers
  * The number of open files
  * Time spent in User, System etc
  * Queues: how many items and how much wait

### Why to use
- help with corelations 
  * Error/time <--> CPU-util/time <--> throughput

### CPU handlers
- Userland (us): your applications
- Nice (ni): Same but altered priorities
- System (sy): Kernel work
- Idel (id): noting ((tricky can be doing something still))
- IOWait (io): time spent waiting disk IO

## ways to chage data
###  1 **gauge**
- lose data from between the time marks
###  2 **Counter**
- `this-happened: this-may-times`
- only loses when exatly it happen between the time marks (which matters very less)
###  2 **Histerogram**
- `this-happened: this-may-times` but in graph way
- Basically a metric vs metric-freqancy graph
### How to use
- we now have the premitive data the System but..
#### how to infer high level data like...
- cluster efficiency
- overloaded node
- core distribution

#### How does it reach the final frontend / Motadata?
##### prometheus
- encodes in utf-8 simple text via http
```prometheus
node_cpu_seconds_total{cpu="0",mode="idel"} 831.26
---
metric_name{key1=value1, ... , keyN=valueN} number
```

##### open Telemetry
- uses [[protobufs]]

### Garbage Collection
- Some metrics are less useful than other metrics
- while exporting these metrics we have to chose which metrics to forward as it is costly to send all of them
- Basically we have to examine the [[ROI]] of the metrics we want to export 


## 🧠 Why This Matters
**Personal relevance:** 
**Broader impact:** 
**Long-term value:** 

## 🔗 Connections & Context
**Builds on:** [[observability#Telemetry|Telemetry]] 
**Relates to:** [[prometheus]] [[grafana]] 

[[pressure-stall-information]]
**Leads to:** [[]] 
**Contradicts/Tensions:** [[]] 

## 📝 Deep Processing Questions
1. **How would I explain this to a 12-year-old?**
   
2. **What's a concrete example or analogy?**
   
3. **What would happen if this weren't true?**
   
4. **How do this change my thinking about X?**

## 🔧 Application & Use Cases
**Immediate applications:**
- 

**Future scenarios where this applies:**
- 

**Action items spawned:**
- [ ] 

## ❓ Knowledge Gaps & Next Steps
**What I don't fully understand yet:**
- 

**Questions to explore:**
- 

**Related concepts to study:**
- 

## 📚 Evidence & Sources
**Primary source:** 
**Supporting evidence:** 
**Counterarguments:** 

---

## 🔄 Review Checkpoints

### First Review (24-48 hours)
- [ ] Can I recall the core concept without looking?
- [ ] Do the connections still make sense?
- [ ] New insights or applications?

### Weekly Review
- [ ] Can I teach this concept to someone else?
- [ ] Have I found new connections?
- [ ] Does this need refinement or expansion?

### Monthly Review
- [ ] Is this concept integrated into my broader knowledge?
- [ ] Have I applied this in practice?
- [ ] Should this be promoted to a higher-level concept note?

---

## 💡 Insights & Evolution
**Initial thoughts:** 

**After first review:** 

**Refined understanding:**
