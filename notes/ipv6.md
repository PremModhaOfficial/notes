---
id: ipv6
aliases:
  - IPV6
tags: []
created: 2025-06-12 18:19:44
modified: 2025-06-12 18:21:47
---

# IPV6

## Problem
- Because we ran-out (of approx 4.5 bill)

## Difference
| v4               | v6                                                             |
| ---------------- | -------------------------------------------------------------- |
| 32 bits          | 128 bits (4 times)                                             |
| 4 octets         | 8 hextets                                                      |
| separated by '.' | separated by ':'                                               |
| has classes      | no classes only IP masks line /64 which is half of the full ip |


## shorten it
- replace continuous zeros with ::
  - Machines replaces this automatically with the missing zeros
  - >[!NOTE] Can Only Have One such locations

- Remove all the leading zeros

```Global_Unicast
 2001:0db8:0000:0000:a111:b222:c333:abcd
 \____________/ |__| \_________________/
 Global Prefix |SubnetID|    IntefaceID|
```










