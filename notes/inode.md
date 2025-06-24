---
id: inode
aliases:
  - inode
  - os/
tags: []
created: 2025-07-01 13:45:47
modified: 2025-07-01 14:01:24
---

# inode
#topic/os/filesystem/inode

i ndoe is a unique number that is related to each file

# what is contains (expect the data and the filename all the metadata is stored)
1. filetype
2. owner and groupID
3. File Size
4. timestamp
5. LinkCount
6. Pointers to data blocks where the actual file data is

the file name in the directory actualy points to this inode number
this helped make saparate metadata from the files


[[symlinks]] has the same inode
