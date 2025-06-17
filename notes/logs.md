---
id: logs
aliases:
  - LOGS
tags:
  - topic/Observability
  - concept/devops
  - source/udemy
created: 2025-06-17 11:29:35
modified: 2025-06-17 13:58:04
status:
  - #draft
type:
  - #atomic-note
---

# LOGS
## 🎯 Core Concept
> [!FAQ] Log 
> The sailors throws a bit of wood attached to a rope containing noughts,
> By counting how many noughts passed to calculate the distance

> [!FAQ] Log For Computers
> Act of keeping/ noting Sequence of event that occur on a computer system

## Logs Components
1. **Context**
  a. On which event did this happen
  b. time of this happening
  c. What is the state of variables in this problem

## Problems
1. Logging Destination
  a. might need its own process and mutex or sync and stuff
2. Improper logging
3. Failure of logging in case of application abrupt malfunction / shutdown
4. Loss of Logs due to memory issues of all kinds leading to abrupt shutdowns of applications causing the loss of data not written to the disk

## Structure
1. To Ensure they are easy to consume (available telemetry but for starters)

## When to use #Logs
now write or wrong answer there are tradeoffs

### GOOD 
  1. Easy To implement
  1. Well supported
  1. Arbitrarily complex

### BAD
  1. Expensive
  2. Unstructured (pattern derivation is hard to do from vanilla logs)

#### But when to use logs
  1. Low traffic
  2. single (limited) Instance
  3. Isolated (best case for logs)
logs are valuable for any type of architecture but it is in rare cases that they do

## Consuming the logs
  - use tools like `jq` and `grep` etc


### exercises
logs implementation in golang: [logs in golang](./code-exaples/logs/main.go)

## Syslog
- A standard for loggin
- before: `/var/log/syslog`
- now: `systemd/journald` backward compatible with syslog


## 🧠 Why This Matters
**Personal relevance:** 
**Broader impact:** 
**Long-term value:** 

## 🔗 Connections & Context
**Builds on:** [[]] 
**Relates to:** [[]] [[]] 
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
