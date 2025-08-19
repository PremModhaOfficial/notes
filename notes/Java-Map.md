---
id: Java-Map
aliases:
  - Map
tags: []
created: 2025-07-29 16:08:47
modified: 2025-07-29 17:10:29
---

# Map

```java
public interface Map<K, V> {
    V put(K key, V value);

    V get(Object key);

    V remove(Object key);

    boolean containsKey(Object key);

    boolean containsValue(Object value);

    int size();

    boolean isEmpty();

    // bulk

    void putAll(Map<? extends K, ? extends V> m);

    void clear();

    // colection view
    public interface Entry<K, V> {
        K getKey();

        K getValue();

        K setValue();
    }

    Set<K> keySet();

    Collection<V> values();

    Set<Map.Entry<K, V>> entrySet();
}

```
## HashMap 
- basic implementation of Map with key hashed and order not preserverd

## LinkedHashMap
### Constructors
```java
public class LinkedHashMap<K, V> implements Map<K, V> {
  public LinkedHashMap() { }
  public LinkedHashMap(int initialCap) { }
  public LinkedHashMap(int initialCap, float loadFactor) { }
  public LinkedHashMap(int initialCap, float loadFactor, boolean accesOrder) { } // LRU cache
}
```

[[cache#LRU|LRU]]

- HashMap but the order is preserved using the  
