---
id: java-queue
aliases:
  - queue
tags:
  - topic/java/collection/queue
created: 2025-07-28 17:59:10
modified: 2025-07-29 10:34:47
---

# queue
- first and Last pointer to facilitate the operations on them
- circular queue as base data structure

- FIFO / LILO

|         | Throws Exception | Returns Special Value |
| ------- | ---------------- | --------------------- |
| Insert  | add(e)           | offer(e) -> false     |
| Remove  | remove()         | poll() -> null        |
| Inspect | element()        | peek() -> null        |



Implemented By: [[Java-BlockingQueue|BlockingQueue]]

## DoubleEnded Que
- [[Java-ArrayDequeue|ArrayDequeue]]
