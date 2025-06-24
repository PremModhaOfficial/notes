---
id: swap-memory
aliases:
  - swap-memory
tags: []
created: 2025-07-07 17:13:03
modified: 2025-07-07 17:28:03
---

# swap-memory

- memory has limits
- [[heap#Vertual Memory]] provide an interfae that can furthur optimize use of memory
- move the unused memory to disk and make space for others
- map this to the Page Table and mark it as not present in the mem and present in the swap



- when the application request the swapped memory
- the mmu triggers the page fault that make the kurnel wake up and throuht the elevated privlages it brings back the swapped memory to the main memory
- and make it the new mapping 
- this is slow and can be felt at user level

zram - compress and keep in memory #topic/os/swap/zram
zswap - compress and swap in disk #topic/os/swap/zswap
OOM kill - just kill the process #topic/os/OOMKILL
