---
id: network-address-translation
aliases:
  - NAT
  - nat
tags: []
---

## The Problem
- We ran out of IPv4 addresses to assign

## Solution
- We can assign same IPs inside an *Internal* area and use a router with a valid and unique IP to 'rout' the traffic
- Inside of this private networks are the private IPs which are reusable

### example
- Inside your house you have the private IPs [[IPV4#Private]] (192 ones)
- They can communicate within this network but currently isolated from the internet
- Now comes the routes with a valid Public IP (unique in the world) (only one is needed usually)



## How
### Types of NAT
### PAT (Port Address Translation) (overload)
 kind of like that image
![[./attachments/images/tci-ip.png]]

- Most common
- Swaps the source IP (previously Device's private IP) IP:port with Router'sIP:prot|next_available_port and sent the packets to net
- When it receives the data back it checks its own table for the mappings and swaps the destination IP likewise
### Dynamic
- Like PAT but the router actually dynamically assigns the device with a public IP from the IP pool purchased from your ISP

> [!NOTE] PORT
> the prot number stays the same while in case of the pat it can chage if the port became unavailable

### Static
- Just Like the other two options
- But you have to define the mappings manually



[[IPV6]]
