---
id: cdp
aliases:
  - CDP
  - cisco discovory protocol
  - CDP cisco discovory protocol
tags: []
created: 2025-06-23 14:16:30
modified: 2025-06-24 13:54:40
---

# CDP cisco discovory protocol

#protocol/Layer2

- Works in cisco devises
- Layer-2 protocol
- find-out what are the capabalities of the neighbours
- to discover what are the neighbours and what thay are capable of?
- A switch/Bridge send messages to the directly connected devises
`Hey I am here `
`What am I`
`My capabalities`
`How to talk to me`

## Works
- multicast to 01:00:0c:cc:cc:cc
- by default every 60 secs if not disabled
- Carries Type-Length-Values TLVs

### TLV
consists of
- DeviceID
- IP address
- Interface Info
- Software Version
- Capabilities (router, Switch, bridge etc)
- platform

> [!note] TLVs are Customizable in case of CDP only

### Analogy
- basicaly a group caht for 2nd layer and introduction channel but only works for neighbours

[[LLDP]]

