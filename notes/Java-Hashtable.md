---
id: Java-Hashtable
aliases:
  - Hashtable
tags: []
created: 2025-07-29 11:13:45
modified: 2025-07-29 11:47:40
---

# Hashtable

- K-V pais 
- Associative array (dictionarry)
- fast O(1) lookups
- uses [[saparate-chaining]] or [[open addressing]]

## Life Cycle
first hash the element
if found empty make new linked list
if collided check if unique and append else return false

> [!note] Override Equals and hashCode methods for any userdefined objects / classes

> [!warning] Resizing 
> Resizing caouses rehashing og the whole tabel

