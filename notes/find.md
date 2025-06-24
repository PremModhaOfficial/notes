---
id: find
aliases:
  - find
tags: []
created: 2025-07-01 12:47:16
modified: 2025-07-01 14:43:08
---

# find
find `opts` `path... ("." by default)` `expr`

## flags
- -name name of file but reqire exact match
  * also used in regex
- -iname name of file but reqire exact match with case insensitivity
  * also used in regex
- -user find file with user ownership
- -inum [[inode]] number

- -size search by the size but reqire exact match
  * SIZES
    + G gb
    + M mb
    + k kb
    + c bytes
  * Rages
    + `-size +1M -size -1G` filtes the results between the 1mb and 1gb files
- -type
  * f for files
  * d for dirs
  * l for links
  * b for block devs
  * s for sockets
- -links find by the number of links the inode contains
- -perm
  * /u=r
  * /r=r
- -newer filename.txt finds the file thah are newer than the provided path
- -empty
- -exec cmd {}
- -mtime DAYS old file

### case: 1: search in the specific path
```sh
find from_path --name named_file
```
### case: 2:
### case: 3:
### case: 4:
### case: 5:
### case: 6:
### case: 7:
### case: 8:
### case: 9:
### case: 10:
### case: 11:
### case: 12:
