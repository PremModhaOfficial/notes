---
id: process-controll-block
aliases:
  - Process-Controll-Block
  - PCB
tags: []
created: 2025-07-07 15:13:10
modified: 2025-07-07 16:06:43
---

# Process-Controll-Block

- The Process has metadata like Execution Context, Registers, Page-Tables, Programme counter, IR etc
- This is managed in PCB that lives in Kernal only Memory #kernel-reserved --> CACHED till L1

- lives in RAM but due to it being shlow it lives on the cache and gets flushed on [[Context-Switch]]

| Category            | Fields/Examples                                       | Links                      |
| ------------------- | ----------------------------------------------------- | -------------------------- |
| **Process ID**      | PID, PPID (parent)                                    | [[pid]]                    |
| **Process State**   | Running, Ready, Blocked, etc.                         | [[Process-Sates]]          |
| **CPU Registers**   | PC (program counter), SP (stack ptr), general-purpose | [[registers]]              |
| **Scheduling Info** | Priority, queue pointers, CPU time used               | [[Scheduling-Information]] |
| **Memory Info**     | [[Page-Table]], segment tables, base/limit registers     | [[memory]]                 |
| **I/O Status**      | Open file descriptors, I/O devices assigned           | [[IO]]                     |
| **Accounting Info** | CPU time used, time limits, user/group ID             |                            |
| **Context Info**    | When switching processes (saved CPU state)            | [[Context]]                |


