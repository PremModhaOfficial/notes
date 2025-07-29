---
id: java-Collections
aliases:
  - Collections
  - Collection(S)
tags: []
created: 2025-07-25 16:59:22
modified: 2025-07-25 17:28:51
---

# Collections
- not to be confused by Collection Interface
- this one here is a class rather than an interface



# Turn Non Synchronized Collections to Synchronized Collections

```java
public static <E> List<E> sychronizedList(List<E> list);
public static <E> Set<E> sychronizedSet(Set<E> set);
public static <K, V> Map<K, V> sychronizedMap(Map<K, V> map);
```
