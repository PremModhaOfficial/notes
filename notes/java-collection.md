---
id: java-collection
aliases:
  - java-collection
  - Java-collection
tags: []
created: 2025-07-25 16:57:31
modified: 2025-07-29 16:08:46
---


# Java-collection
# Collection

## To Rrefer a group of Object as a single entitiy

## Hierarchy
[[collection.excalidraw]]

implements Serializable and cloneable

>[!note] Serialize?
> Collection Extend Serializable interface that means  every collection is serializable
> and cloneable


## methods

### ADD
```java
boolean add(Object);
```

```java
boolean addAll(Collection);
```

### Delete

```java
void clear(Collection);
```

```java
void remove(Object);
```

```java
void removeAll(Collection);
```

```java
void retainAll(Collection);
```

### Info?
```java
boolean isEmpty();
```

```java
int size();
```

```java
int contains(Object);
```

```java
int containsAll(Collection);
```

```java
int iterator();
```

```java
int toArray();
```

## Also classified
### basic
- add
- remove
- size
- isEmpty
- iterator
- contains
### bulk
- addAll
- removeAll
- retainAll
- containsAll
- clear
### Array
- toArray() -> Object[]
- toArray(T[] t) -> T[] // if t can it will be filled othewise it will be replaced with the sutable new one
 


Implemented By:
[[java-List]]
[[java-Set|Set]]
[[java-queue|queue]]


[[Java-Map|Map]]
[[Java-Hashtable|Hashtable]]



