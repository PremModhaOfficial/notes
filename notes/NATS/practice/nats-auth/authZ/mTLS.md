---
id: mTLS
aliases: []
tags:
  - #nats
backlinks: []
created: 2026-01-12 12:24:06
links: []
modified: 2026-01-12 14:17:21
status: reviewed
---

# why
- the traditional tls assumes the server is trusted;
- but in case of other architectures (like micro-servises) sometimes there is no concrete server and client

# what
- the server and clinet both now hase similar hierarchical position and will both need to have certificates of theire own
- then they both mutualy authenticate each-other and proceed so to enhance security and reduce rics of man in the middle attacks

# HOW
- first client authenticates with the servers as normaly done before
- then the server will also authenticate the server
- basically gives you trust in no-trust environment
