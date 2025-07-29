---
id: java-list
aliases:
  - java-List
  - List-Interface
  - java-List-interface
tags: []
created: 2025-07-25 13:43:36
modified: 2025-07-29 09:58:17
---

# java-List-interface
#topic/java/interface/list
implements: [[java-collection|Collection]]

## Why
- when you want preserve the sequesce of insertion
- when you want to allow duplicates

## How are we going to do that??
- Make use of Indexes

## Overrides


## Overloads
```java
boolean add(int index, E e);
boolean addAll(int index, Collection<? extends E> c);
E remove(int index);
```

## NEW

```java
E get(index i);

// @throws IndexOutOfBoundsException - if the index is out of range (index < 0 || index >= size())
E Set(index i, E e); // returns the previosue Element at this index

// -1 if not found or gives the first accurance
int indexOf(E e);
int lastIndexOf(E e);

// ListIterrator
ListIterator listIterrator();
```


> [!note] Arraylist and Vector
> Implements Random Access interface that basicaly means index based access which is O(1) for any index

### Also classified as 
#### Positional
- add(int index,E e);
- addAll(int index,Collection e);
- set(int index,E e);
- remove(int index);
#### Search
- indexOf(E e)
- LastIndexOf(E e)
#### Iteration
- iterator()
- listIterator()
#### Range-view
- subList(int from, int to) backed by original 


Implemented By :
  - [[java-ArrayList|ArrayList]]
  - [[java-LinkedList|LinkedList]]
  - [[Java-Vector|Vector]]
  - [[Java-Stack|Stack]]
