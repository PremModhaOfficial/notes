---
id: ddr-dram
aliases:
  - DDR-DRAM
tags: []
created: 2025-07-07 16:55:45
modified: 2025-07-07 17:01:53
---

# DDR-DRAM

double data rate [[DRAM]]
- read/write on UP and DOWN tick of the clock
- double the rate at witch you do io

# DDR5
- the ddr5 has made it so it hase 16 io pins so it can be split into two byte channel one byte each
  * this helps in parallel memory access


# Internals
## Banks
- banks are connected to a collection of rows 
- banks only exposes one row to the bus at a time
### ROW
- the row has (1 cache line worth of data in collumns)
#### Collumn
- houses the capasitor-transistor pair - a bit
