---
id: ospf-multiarea
aliases: []
tags: []
---


## Divide the network in saparaete areas
-  The subnets are Divided and contained within, and communicated via the connecting router
-  The updates and changes are taken care of inside the said areas and the other areas and Backbone area don't have to know about it.

- Terminologies
  -  Routers Types Based on Their Location in Areas
    - All the routers inside or connected to the Backbone area are Backbone Routers
    - All the connecting router are known as area border routers (ABR)
    - All the router inside an area except the backbone is Internal router 
    - All the router that connects to other no ospf routers is (ASBR) (ex. connected to internet)

[[OSPF Multiarea.canvas|OSPF Multiarea]]


[[BGP|next]]
