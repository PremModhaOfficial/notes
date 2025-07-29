---
id: Java-Vector
aliases:
  - Vector
tags: []
created: 2025-07-28 11:13:01
modified: 2025-07-29 09:59:03
---

# Vector

- same as [[java-ArrayList|ArrayList]] but synchronized and a legacy class
- Also has the specific method becouse of being a legacy class the method names are verbose like addElement removeAllElement etc

## Specific Methods
```java
void addElement(E e); //vector
void add(int index, E e); // list just fail if not possible
boolean add(E e); // vector



boolean remove(E e); // Collection
boolean removeElement(E e); // Vector
E remove(int index); // List
void removeElementAt(int index) // vector
void clear(); // collection
void removeAllElement() // vector
```
