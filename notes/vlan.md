---
id: vlan
aliases:
  - Vlan
  - VLAN
tags: []
created: 2025-06-20 17:56:34
modified: 2025-06-20 18:33:40
---


# VLAN
- Vertual Lan
- actualy a single switch which can be logicaly partisioned to act as 2 or more lans / broadcast domain
- partisioned at layer 2
- vlan is in diffrent broadcast domain
- 2 vlan need routers to comunicate
- host in a vlan are unaware of it being in vlan

# Benifits
- security
- cost reduction
- better performace
- shink broadcast domain
- improved IT staff effic.

# Type of VLAN
## static
- based on port numbers, called port based vlan
- manually assigned
- a port exist in only a single vlan

## dynamic
- based on mac address
- automatically assigned by the switch
- called mac based vlan
- dynamic vlan config require the VMPS (vlan membership policy server)

# link types
## access link
- no vlan information
## trunk link
- the vlan info is included
# How they are identified
- *By frame tagging* #frame_tagging
- Trunk Link is the link that comes to the vlan 
- it has to have vlan identifier to tell vlan where to transmit 
- it strips this metadata when it forwards to the connected host
- also inserts when it  comunicates with the trunk

## Trunking protocols
### ISL
- Inter Switch Link
- cisco proprietary protocol
- Supports ethernet token-ring FDDI
- Adds 26 bit header and 4 bit trailer containing crc to ensure data integrity

### IEEE 802.1Q
- only adds 4bit of tag

[[stp|Spanning Tree Protocol]]
