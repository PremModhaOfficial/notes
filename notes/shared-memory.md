---
id: shared-memory
aliases:
  - Shared-Memory
tags: []
created: 2025-07-07 17:04:41
modified: 2025-07-07 17:06:46
---

# Shared-Memory
- if multiple processes/threads have same memory the mmu tries to not have copy and just refrences that page instead
and uses [[COW]] for better optimization
