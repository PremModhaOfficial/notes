---
id: java-navigable-set
aliases:
  - Java-Navigable-Set
tags: []
created: 2025-07-29 13:55:56
modified: 2025-07-29 14:32:46
---

# Java-Navigable-Set
```java
interface NavigableSet<E> extends Set<E> {

    E lower(E e);   // greates < E
    E floor(E e);   // greates <= E
    E ceiling(E e); // least >= E
    E higher(E e);  // least > E
    E pollFirst();
    E pollLast();

    NavigableSet<E> descendingSet();
    NavigableSet<E> subSet(E fromElement, boolean fromInclusive, E toElement, boolean toInclusive);
    NavigableSet<E> headSet(E toElement, boolean inclusive);
    NavigableSet<E> tailSet(E fromElement, boolean inclusive);
}

```


Emplemented By [[Java-TreeSet|TreeSet]]
