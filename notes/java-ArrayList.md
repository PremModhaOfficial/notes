---
id: java-ArrayList
aliases:
  - ArrayList
tags: []
created: 2025-07-25 16:02:28
modified: 2025-07-28 11:22:23
---

# ArrayList
#topic/java/interface/list/ArrayList

## when to chose this??
- when frequest operation is retival `O(1)`

- java.util.ArrayList;

## What is so Special?
- It is a concrete Object so has Constructors
- uses good'ol Array in the backend and dynamicaly handels its size
- Stores duplicate
- holds Heterogeneous data (Only treeSet and treeMap has homogeneous data in whole of collection)
- Has RandomAccess with indies  (only ArrayList and Vector implements RandomAccess interface)

```java
ArrayList list_default_capCT = new ArrayList(); // with 10 as capacity
ArrayList list_init_capacity = new ArrayList(int initialCapacity); // with specified capacity
ArrayList list_with_collecTN = new ArrayList(Collection<? extends E> c); // with objects of the previouse Collection c
```

## Capacity
- dynamicaly increses the size by making it (3n/2 + 1) where the n is last capacity
- it start with 10 as default then goes to 10 -> 16 -> 25 -> 38 -> ...


[[ArrayList-VS-Vector]]
