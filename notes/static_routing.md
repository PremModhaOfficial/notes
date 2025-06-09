---
id: static_routing
aliases:
  - Static Routing
tags: []
---


# Static Routing
 - manual 
 - admin decides the path 
 - Mandatory need for destination Network ID
 - IS secure and fast
 - Has 0 or 1 Administrative The lower the better
 - Disadvantages 
  - Only in small orgs (10-15 routs)
  - manual
  - network change effects whole network

## HOW TO
```
Router(config)#

ip_route <destination network ID> <destination subnet mask> <Next-hop IP Address>
```

