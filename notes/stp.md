---
id: stp
aliases:
  - Spanning Tree Protocol
  - Redundacy is good
tags: []
---

# Redundacy is good
 *Problem*
   - To prevent total loss of cimunication we add redundant links in our networks.
   - but because of redundant connections a broadcast may find devicce with multiple ways
   - or worse may find a circular path in witch it keeps repeating (#braodcast-storm) and networks goes down
# Spanning Tree Protocol
*Solution*
 - we just mark the redundant link as backup link and only use incase of main link failear
## How To do that
- sitches probe the networks by sending messages named probe to discover loops
### Probe (message)
 - has [[BPDU]] (Bridge Protocol Data Unit) which contains specific data about the sender switch 
 - evry switch multicasts its oun bpdu, when switch encounters message with the same bpdu as it self it discoers the loop
 - this happens every 2 seconds
 - this is how the root Bridge is elected 
 - all the switches find best way (path) to reach the root bridge
 - all other paths are marked as redundant and will be avoided untill necessary

#### Root Bridge Election
  - one with the higest Priority (Every Switch has MAX_VALUE=32768 by default) OR lowest mac address
  - this make it tie by default and macs are compared and older switches win
  - One with the lowest bridge_id is the root bridge (bridge_id=Bridge-Priority.concat(MAC))
  - why have the Bridge-Priority High By default?
    because newer devices have better features and performance but has higher MAC thus you can manualy chage the root device by lowering Bridge-Priority
[[port_roles]]
[[routing]]
