---
id: BGP
aliases:
  - show ip bgp summary
tags: []
---

## [[autonomous-systems|Autonomous Systems]]

## [[IGP]] & [[EGP]]

## basic-bgp-features
- open standard
- Application layer protocol (TCP 179)
- exterior Gateway protocol
- for Inter-AS domain Routing
- Designed for scalable systems like internet
- Updates are incremental and triggers
- Path vector protocol
  - Carries the paths (ASN numbers it has passed through)
    - helps remove loops
### bgp-features(contd)
- Sends the updates to manually defined Neighbours via unicast
- Metric = Attributes
- #admin_distance 
  - 20 external updates EBGP
  - 200 internal updates IBGP

## bgp-loop-prevention mechanism
- Prevent the connection if my own ASN is in the List

## BGP neighbours
- These are the routers forming a tcp connection for BGP updates
- Also known as BGP Peers and BGP speakers
### Types of BGP different from [[IGP]] and [[EGP]]
- IBGP (Internal)
  Routers are IBGP Neighbours if they are from *different* ASN

- EBGP (External)
  Routers are IBGP Neighbours if they are from *same* ASN

## DB
### Just like [ospf](notes/ospf.md#how-it-happens) #LSDB it has a database
  - Contains the list of all configured neighbours.
```cisco ios
# show ip bgp summary
# show ip bgp neighbours
```
### Forwarding Table/Database
- List of networks known by the BGP with their paths and attributes

```cisco ios
# show ip bgp
```
### IP Routes
List Of best paths to destination networks

```cisco ios
# sh ip route
# show ip route
```

[[virtual-router-redundancy-protocol]]
