---
id: port_roles
aliases:
  - port roles
  - port types
  - Port Roles
tags: []
created: 2025-06-20 18:43:18
modified: 2025-06-20 18:57:55
---



# Port Roles
## Root Port (RP)
  - To reach The Root bridge
  - Each switch has only one of them
  - Has the lowest cost to reach the root
  - root bridge adverties the [[bpdu]] with lowest cost other bridges add cost of the reciever port to that bpdu and forward throug other ports
## Designated Port (DP)
  - Forwarding port
  - One per Link
  - also lected with lowest cost
  - every switch must have one regardless of root port existance

## Election Of ports in order
1. lowest path cost to root
2. lowest bridge id
3. lowest sender [[port-id]]


## Blocking/Non-Designated Port (Loops)
  - loop

## why
We needed the labeling to identify loops and we have space for other links so use defined other roles

## PORT ID

