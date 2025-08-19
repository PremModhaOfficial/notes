---
id: Java-Comparator
aliases:
  - Comparator
tags: []
created: 2025-07-29 14:26:10
modified: 2025-07-29 14:30:32
---

# Comparator
- interface
- Compares two objects the custom way 
- used in TreeSet and TreeMaps to provide custom objects with custom Sorting // needed by the set to have to compared for sorting


```java
public interface Comparator<E> {
  int compareTo(E e) 
    /*
     * returns -1 if e is bigger and need to go to left
     * returns  0 if e duplicate
     * returns  1 if e is smaller and need to go to right
     */

}
```
