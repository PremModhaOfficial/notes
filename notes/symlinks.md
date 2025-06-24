---
id: symlinks
aliases:
  - symlinks
tags: []
created: 2025-07-01 13:53:03
modified: 2025-07-01 13:59:06
---

# symlinks

## HARD links
- hard link pionts to the same inode
- delete one inode and by extention the actual data still reamins it just decrements the link count
- when counter reaches zero it deletes the actual data
## SOFT links
- #pending


## Rename
when you move/reanme the file it only upadates the directory entries because the actual inode remains the same

## Create (file cration)
- Create the inode first and alocate the memory
- then link it with the actual file path

