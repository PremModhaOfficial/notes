---
id: heap
aliases:
  - Heap
tags: []
created: 2025-07-07 15:36:24
modified: 2025-07-07 17:30:05
---

# Heap

- the memory boundry growing opposite of the stack

## Storage
- memory is stored with an abstraction
- Logical Memory Address -> physical Memory Address
- This abstraction is provided by the OS Kernel and MMU

## Vertual Memory
- The raw memory management is prone to security exploits and mistakes like [[Dangling-Pointer]] and [[Memory-Fragmentation]]
- The os provides an abstraction that help programme declare the memory related oprations
  * var = maloc(type, size);
  * free(vak);


### Internals of vertual Memory
- The memory is stored in whats called the pages sized at power of 2 like 4KB, 16KB, 64KB
- This block is alocated to the process and kept track of by the [[Page-Table]] which is kept in the Memory
- this entries are cached in something called TLB -[[Translation-Lookaside-Buffer]] once per core

