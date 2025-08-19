---
id: java-sortedset
aliases:
  - Java-SortedSet
tags: []
created: 2025-07-29 13:55:54
modified: 2025-07-29 14:05:02
---

# Java-SortedSet

## Interface Overview
```java
public interface SortedSet<E> extends Set<E> {
    // range view
    SortedSet subSet(E from, E to);
    SortedSet headSet(E toElement);
    SortedSet tailSet(E toElement);

    // Endpoints
    E first();
    E last();

    // Comparator Access
    Comparator<? super E> comparator();
    default Spliterator<E> spliterator() {}
}
```




Implemented by : [[java-navigable-set]]
