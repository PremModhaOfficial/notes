---
id: ospf
aliases:
  - OPEN Shortest Path First
tags: []
---


# OPEN Shortest Path First

## Why is ospf
- To learn the routs 

## what is ospf
- Routing protocol widely available in almost every router
- Interior gateway protocol (IGP)
- Link State routing protocol

## How it happens
- Let every router learn about every other router
  - By using Link-State Advertisement (LSA) to every router
- Now every Router has an LSDB (LS-DB) 
- For this to work you must have the same database on every router

## Deeper Depth
### DOWN State
- Become Neighbours
  1. Choose a router id (RID) (IPv4)
  2. IP assignment
    - Manual assigned
    - IF NOT Router is assigned the highest up - status loopback interface IP
    - IF NOT Router is assigned the highest up - status non-loopback interface IP
  3. Sends HELLO my RID is 1.1.1.1 Neighbours: \[\]
  4. Other router revives this message and checks for these strict requirements
    - Area ID
    - Subnet
    - Hello and Dead interval
    - Authentication
    - Stub area flag
    - Unique router id
- Share LSDB info
- Choose Best routs

### INTI State
- Router 2 is now in init stage
- Send the message with Neighbours: \[1.1.1.1\] to 1
- 1 gets the message and gets promoted to 2-way stage
- now router 1 also send the message with know Neighbours: \[2.2.2.2\]


### Problem: in case of change in multipoint broadcast-storm can occur
- *Solution*: make a Designated routers (DR) and backup-DR (BDR) using automated election using router id and priorities

[[ospf-multiarea]]
