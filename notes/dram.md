---
id: dram
aliases:
  - DRAM
  - Dynamic RAM
tags: []
created: 2025-07-07 16:45:57
modified: 2025-07-07 16:52:16
---

# DRAM
- dynamic Ram
- 1 transitor and capasitor per bit
- relatively cheper than the [[SRAM]] but slower
- needs to refresh the capasitor to avoid the curruption of data
- this is why it cannot be read cocurrently directly through the direct access

# problem 1
- needs constanst refresh of bits becouse capasitors dischage over time

## Solution  read/write throgh cache
- when wnat to read take the cache line / row of the data and write it to the cache and read it from there
- when done or in need of some space flush the cache and re write the data
