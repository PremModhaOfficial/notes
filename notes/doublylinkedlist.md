---
id: doublylinkedlist
aliases:
  - DoublyLinkedList
tags: []
created: 2025-07-25 17:33:30
modified: 2025-07-28 10:46:42
---

# DoublyLinkedList

```
+-------+     +-------+     +-------+
| Head  | --> | Node 1| --> | Node 2| --> ...
+-------+     +-------+     +-------+
    ^             ^             ^
    |             |             |
   null         Data1         Data2
    |             |             |
    v             v             v
+-------+ <-- | Prev  | <-- | Prev  | <-- ...
|       |     +-------+     +-------+
| Tail  |
+-------+
```

**Lifecycle of a Doubly Linked List (Illustrative Steps):**

1.  **Creation:**

    *   An empty doubly linked list is created. `Head` and `Tail` pointers are both `null`.

    ```
    Head: null
    Tail: null
    ```
2.  **Insertion (First Node):**

    *   A new node is inserted. This node becomes both the `Head` and the `Tail`. The `Prev` and `Next` pointers of this node are `null`.

    ```
    +-------+
    | Head  | --> +-------+
    +-------+     | Node 1|
                      +-------+
        ^             ^
        |             |
       null         Data1
        |             |
        v             v
                    +-------+ <-- null
                    | Prev  |
                    +-------+
    +-------+ <--
    | Tail  |
    +-------+
    ```
3.  **Insertion (Subsequent Node):**

    *   A new node is inserted after the `Head`. The `Next` pointer of the `Head` points to the new node. The `Prev` pointer of the new node points to the `Head`. The `Tail` pointer is updated if the new node is inserted at the end.

    ```
    +-------+     +-------+     +-------+
    | Head  | --> | Node 1| --> | Node 2|
    +-------+     +-------+     +-------+
        ^             ^             ^
        |             |             |
       null         Data1         Data2
        |             |             |
        v             v             v
                    +-------+ <-- | Prev  | <-- +-------+
                    |       |     +-------+     |       |
    +-------+ <------- SEARCH
 | Tail  |                   |       |
    |       |       +-------+                   |       |
    +-------+                                   +-------+
    ```
4.  **Deletion:**

    *   When a node is deleted, the `Next` pointer of the previous node is updated to point to the next node of the deleted node, and the `Prev` pointer of the next node is updated to point to the previous node of the deleted node. If the deleted node is the `Head` or `Tail`, the `Head` or `Tail` pointer is updated accordingly.

5.  **Traversal:**

    *   Traversal can occur in both directions, forward (from `Head` to `Tail`) and backward (from `Tail` to `Head`), using the `Next` and `Prev` pointers, respectively.

