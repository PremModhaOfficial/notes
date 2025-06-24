---
id: memory-fragmentation
aliases:
  - Fragmentation
  - Memory-Fragmentation
  - Memory-Management-Unit
  - MMU
tags:
  - #topic/os/fragmentation
  - #topic/os/memory
created: 2025-07-07 16:06:53
modified: 2025-07-07 16:42:10
---

# Memory-Fragmentation
- the sate of having left with  non-contiguous memory
- this hinders the processes becouse it can be used in diffrent place

## Internal Fragmentation
- when the os alocates fix sized memory to the processes we are left with unused space, which may add up to substantial ammount and may block a potential process

## External Fragmentation
- Lets alocate based on their need only for efficiency.
- But when we alocate the memory to process the processes may complete at diffrent time and the alocated memory is freed but the new process will reqire diffrent ammount of space that may or may not be efficently allocatable

## Solution
We made a middle man Memory Management unit that handels this by
- alocating on behalf of the process
- alocating parts of the memory on the fragmented holes and keeping a track of it using [[Page-Table]]
