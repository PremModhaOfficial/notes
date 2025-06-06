---
id: network-devices
aliases:
  - diffrent devices works with diffrent layers in #osi model
tags:
  - osi
---

# diffrent devices works with diffrent layers in #osi model

## repeater
- 2 port device that works on physical layer
repeats the signal from a node to another node
helps fight attenuation over extended transmission

## Hub
- multiport repeater
has 2/4/8 connections

it is not an inteligent device
it will forward the data to every thing hub is connected to


## Bridge
- Data link layer 
- 2 ports
- joins 2 LANS within same subnet (*Must be within same subnet*)

## Switch 
- can have 8/12/16/24
- Data link layer 
- multiport bridge 
- joins two subnets 
- diffrent subnests needs a router to cumunicate with each-other 
### Special Switch 
works on layer 3 and works like a router (making it a router that also does the work of a switc)

Bridge - cam table - software
Switches - cam table - Hardware

## Router
- inteligent device that does not flud the data to all the connected device
- connects the diffrent subnets on layer 3 using ip addresses routers the data to Destination

