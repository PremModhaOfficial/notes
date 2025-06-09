---
id: default_routing
aliases:
  - DEFAULT ROUTING
tags: []
---


# DEFAULT ROUTING
- send all the reqiest to a single fixed gateway

*Represented as* 0.0.0.0/0 (ipv4) , ::/0 (ipv6)

- AD #admin_distance is 255 (lowest priority in Routing Tables)
- Handels traffic for unknown destinations
- these are not fault tolerant
- reqires less resoureses and bandwidth
- no load ballancing
