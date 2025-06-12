---
id: collision-and-broadcast-domains
aliases:
  - collision domain
tags: []
---


# collision domain
An area in which the collision of packets can take place while using a shared channel

When you connect via hubs and 2 or more devices send message in case of half-duplex it will collide and the corresponding area/link is called the collision domain.

- This means that every device connected via hub is inside a same collision domain.
- And every devices connected via switch is in a separate collision domain.

# broadcast domain
An area in which all the divines will receive the broadcast message

- Broadcast domain are separated by routers


> [!FAQ] DOMAIN COUNTS
> AN intelligent device separates the broadcast domain (switch is intelligent device) (counted separably even if it is the same one for switch and router)
> each port of the switch and router is a separate collision domain
