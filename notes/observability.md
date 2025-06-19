---
id: observability
aliases:
  - Observability
  - observability
tags:
  - topic/dev-ops
  - concept/observability
  - source/udemy
  - telemetry
created: 2025-06-16 18:06:42
modified: 2025-06-18 10:24:26
status:
  - #draft
type:
  - #atomic-note
---

# observability
## 🎯 Core Concept #trait
### What does it mean be observable?
> [!NOTE] Definition
> A measure of how well can a system's internal status can be inferred
> from its external outputs


# Telemetry
> [!Note] telemetry
> the automated process of measuring and transmitting data from 
> remote or inaccessible sources to a central location for monitoring and analysis.
> #telemetry

> This data is typically collected using sensors and other devices and then transmitted via communication systems 
> to a central system for processing and interpretation.


> [!Note] Definition (Personal)
> how well can you reason about the internal states of a system by using its telemetric outputs

### Why it is needed?
- 'Modern' software architecture is become very complex and has need of central entity for operations
- So you need to get the data in a central location to debug and damage control
- The Service is very complex set of layers working with each other and influencing the layers above
  1. Business logic
  2. Application Logic
  3. Runtime
  4. OS / Kernel
  5. Hardware

### Kinds Problems
- In solving A problem we need the information and there is [[problems-paradims|4 types of information]] 
  1. KK -> Server responds 500
  2. KU -> Disk writes
    - We don't know information on disk I/O no #telemetry
  3. UU -> we don't know what we don't know

## 📝 Deep Processing Questions
1. **How would I explain this to a 12-year-old?**
- in IT the complex working application or a service depends on layers of technologies.
- When something doesn't work you need to find in which layer the problem has occurred
- If you don't know on which layer the problem lies you have to do much trial and error to find the actual fault
- At scale this can lead to loss of time and money
- Thus we need a system that can tell you the problem and allows you to OBSERVE the systems faster and efficiently
   
2. **What's a concrete example or analogy?**
   
3. **What would happen if this weren't true?**
   
4. **How did this change my thinking about X?**

## 🔗 Connections & Context
<!-- **Relates to:** [[]] [[]]  -->
**Leads to:** [[]] 
**Contradicts/Tensions:** [[]] 
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
