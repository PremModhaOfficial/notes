---
id: lsblk
aliases:
  - lsblk
tags: []
created: 2025-06-30 18:12:13
modified: 2025-07-02 15:23:38
---

# lsblk

list block devises information


## Flags

- b display sizes in bites
- h human readable sizes
- n not to print headers
- o \[format\]
- m MODE
- f metadata like uuids
- I include by major number of the interface > [!NOTE] Will print the childerns of thoes major numbered bloks containing other than the major number specified
- e exclude by minor number
- s reverser the depencdacies
- -P --pairs actualy key-value pair of the information rather than the tabled format 
- -p path that is complete path of the device (works in the lvm)
