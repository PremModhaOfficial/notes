---
id: dependency-inverstion
aliases:
  - dependency-inverstion
tags: []
created: 2025-06-20 16:42:34
modified: 2025-06-27 16:34:29
---

# dependency-inverstion
## Motivation
- Aplicatins directly dependent on the low level moduels causes problems when the underliying software changes
- like database mail service module if the method name chages the whole service mach collapse

- chages in the low level moduels will cause the wave of chages in whole dependency chain


## Solution
- #designepatterns/adapter
- use a middle layer of abstraction that both sides adheres to

> [!faq] Inversion Of Controll
> Before the dependent class creates the dependency inside it self
> After the dependency is injected inside the Dependent class with compile time dynamic dispach



## real life examples
[[doom]]
