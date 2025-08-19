---
id: java-LinkedList
aliases:
  - LinkedList
tags:
  - Queue
  - LinkeList
  - List
created: 2025-07-25 17:29:25
modified: 2025-07-29 10:32:27
---

# LinkedList
## what
- doubly LinkedList as underlying data-structure > [[doublylinkedlist|doubly-linked-list]]
- Not implementing Random Access (dont have the Index based fast lookup but emulates the index)
- can have duplicates

## Methods
```java
void addFirst(E e)
void addLast(E e)

E getFirst(); // from the Queue interface
E getLast();  // from the Queue interface

E removeFirst(); // from the Queue interface
E removeLast();  // from the Queue interface

E get(int index);
int add(i)

void remove(E e)
indexOf(E e)
lastIndexOf(E e)

```
## Construction

```java
new LinkedList();
new LinkedList(Collection c);
```

[[arraylist-vs-linkedlist]]







