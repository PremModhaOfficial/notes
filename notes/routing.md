---
id: routing
aliases:
  - What is routing
tags: []
created: 2025-06-23 14:15:43
modified: 2025-06-26 14:22:37
---


# What is routing
  - the prosses of forwarding packets from and to diffrent network
  - and finding the best path to rout the packets
# What is a route
- for a router it is like this 
  * if i want to send the packet to destination X
  * i need to send the packet(s) to Y
  * The Y is the next-hop of the router for the specific destination x and the specific router it self

## Types Of Routing
- [[static_routing]]
- [[default_routing]]
- Dynamic (most common)
 - automatic 
 - routers decides the path 
 
> [!note] Automatic Routs
> When we enable a interface by typing `no shutdown` two routs are automaticaly added to the routers routing table
> One of it is the ip address with aproptiate mask like any other ip with a zero where the hosts are supposed to be
> while the oher one is the one with the self's ip with the mask of 32
> This means that if the [[ip_match]] with the second device means that
 
[[cdp]]

[[routing-protocols]]
