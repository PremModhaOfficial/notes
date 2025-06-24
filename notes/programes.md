---
id: programes
aliases:
  - Programes
  - Programe
tags: []
created: 2025-07-02 18:17:12
modified: 2025-07-07 15:36:21
---

# Programes

- code is [[compile|compiled]] and [[linking|linked]] for a cpu to use
- [[process-controll-block|PCB]] is maintained in cache and will be flushed at context switch

```ascii
Process Management in OS
+----------------+
|    Stack       |
+----------------+
|  Kernal Space  | #kernel-reserved
|    Reserved    |
+----------------+
|                |
|    Grows ↓     |
+----------------+
         ↑
         |
         ↓
+----------------+
|    Grows ↑     |
|                |
|    [[Heap]]    |
+----------------+
| Data + Static  |
+----------------+
|  Text/Code     |
+----------------+
```




