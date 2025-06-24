---
id: ssh
aliases:
  - ssh
tags:
  - topic/ssh
  - concept/os
created: 2025-06-24 17:17:18
modified: 2025-06-24 18:32:53
status:
  - #draft
type:
  - #atomic-note
---

# ssh
## 🎯 Core Concept
- Provide Secure connection in insecure network

## Story
- client request the server for connection
- Server responds with the public key
- client checks for its entru in known-hosts file

### layers
- User Authentication Layer
- ssh is an application layer protocol that is based on tcp
- Transport Layers

there are three encription used in a ssh session lifetime
#### Establish Connection- Key exchage (public) - Authentication
![[ASymmetric-encryption]]
once the Authentication is done, both the side compute a shared secet key called session key 
#### Data Transfer
![[Symmetric-encryption]]
- session key is a symmetric key that now is used for the entire session to encrypt the data Transfer
#### Hashing


## 🧠 Why This Matters
**Personal relevance:** 
**Broader impact:** 
**Long-term value:** 

## 🔗 Connections & Context
**Relates to:** [[networking]] [[]] 
**Leads to:** [[]] 

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
