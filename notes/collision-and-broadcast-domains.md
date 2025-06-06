---
id: collision-and-broadcast-domains
aliases:
  - collision domain
tags: []
---


# collision domain
an area in which the collision of packets can take place while using a shared channel

when you connect via hubs and 2 or more devices send message in case of half-duplex it will colide and the coruspondin area/link is called the collision domain.

- this means that every device connected via hub is inside a same collision domain.
- and every devices connected via switch is in a saparate collision domain.

# braodcast domain
an area in which all the divises will recieve the broadcast message

- broadcast domain are saparated by routers


> [!FAQ] DOMAIN COUNTS
> A inteligent device saparates the broadcast domain  (swich is inteligent device) (counted saparatly even if it is the same one for swich and router)
> each port of the switch and router is a saparate collision domain
