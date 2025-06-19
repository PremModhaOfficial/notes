---
id: span
aliases:
  - Span
  - spans
tags:
  - topic/traces
  - concept/observability
  - source/udemy
created: 2025-06-17 18:24:37
modified: 2025-06-18 14:53:33
status:
  - #draft
type:
  - #atomic-note
---

# Span
## 🎯 Core Concept

- Reprecent a single oparation within trace
```TRACE
START:
END:
NAME:
```

## where have i Seen this
   1. in [[encore]] where they showd me the Lifecycle of the request
     a. how much time the request took to prosses and breaks it into db query and waiting etc

## Span Data
- `Span.kind = Internal`
- `Span.attribute = some-attrib`
  1. Attributes
  2. Events
  3. Links
  4. Status
  5. Kind

## Span Kinds
- Server
- Client
- Producer
- Consumer
- Internal (default)

## Span Status
- Error
- Ok

## Span Attributes
- Key-Value pari to atach data to a [[span]] or [[metric]]
- The span kind and status were attribute once but now has become first-class citizens


## Span Events
- Events allows traces to be like logs and was know as logs before

## TOO MUCH!!?
- All this namespaced variables and events and metrics are hard to inplement inside the codebase
- it may result in more telemetry logic than business-logic
- #opentelemetry based sdks exist and give you already made performant abstractions called [[instrumentation]]


## Span Context
1. spanID
  - 8-byte array with atleast one non zero bytes
2. TraceID
  - 16-byte array with atleast one non zero bytes
3. TraceFlag
  - current detail about the trace, ==If it is sampled==
  - If it is sampled
4. TraceState
  - Vendor Specific
5. IsRemote
  - If the span recieved from somewhere else

## Context propagation

> [!faq] Definition
> A propogation machanism that carries execution-scoped values accros api boundries and logicaly associated execution ubits.
> ---
> - opentelemetry

- context object it self is provided by languages Specific ways/features
## go
- 
```go
import (
	"context"
)
func ex() {
  ctx := context.Background()
}
```

## java 
```java
class ThreadLocal<T> {}
```

## propogators #context_propogation
- The propogation in NOT the same for http and event-stream
- Generaly you have a middleware on handler functions on urls which handels the logic for the propogation of trace context

## Baggage
- just span context of other spans
- helps with corelation within trace and spans while in the prosses of genarating span
### consideration
- Baggage size
- Secret information


## Tracing Overhead Mitigation
f

| Technique                   | Purpose                                        |
| --------------------------- | ---------------------------------------------- |
| **Sampling**                | Collect only a fraction of traces (e.g., 1%)   |
| **Batch Exporting**         | Send data in batches, off main thread          |
| **Limit Span Detail**       | Don’t capture unnecessary attributes or events |
| **Exporter Throttling**     | Rate limit outgoing trace data                 |
| **Asynchronous Processing** | Avoid blocking request threads                 |

## 🧠 Why This Matters
**Personal relevance:** 
**Broader impact:** 
**Long-term value:** 

## 🔗 Connections & Context
**Builds on:** [[]] 
**Relates to:** [[]] [[]] 
**Leads to:** [[traces]] 
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
