---
id: java-lifetime
aliases:
  - Type LifeTime
tags: []
created: 2025-07-10 14:40:06
modified: 2025-07-10 15:04:55
---

# Type LifeTime

1. load the class
2. verify that it is not malicious
3. allocate the memory (only static variable and static accessed variables may be alocated at this time)
4. load refrenced classes
5. execute the code

*Normaly No interface is loaded first but in mordern java if interface have a static variable inside it may be loaded*

## Loading Type


## When a class is loaded??
