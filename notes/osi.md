---
id: 1749186058-osi
aliases:
  - osi
  - OSI
tags: []
created: 2025-06-16 15:17:00
modified: 2025-06-16 15:17:01
---

# OSI

- [ ] It is a reference model for machines to communicate via Electronic means
some machine may not exactly mimic the model
all of these layers talk to the layers adjacent to them

All layers wrap and unwrap the data breaks into packets

![[./attachments/images/tci-ip.png]]

# modern TCP/IP
## Application
Http Protocol lives here

talks between applications

> [!FAQ] DATA
> HTTP Request
> GET https://getdata?page=1

## Presentation
- Encrypt and decrypt the data
## Session
- manages sessions
## Transport


Port Numbers To track sessions
provides host-to-host communication
> [!FAQ] DATA (broken into Segments)
> TCP
> Source-Port = 80
> Destination-Port = 7268
> HTTP Request
> GET https://getdata?page=1

## Network
IP addresses
- Provides connection between host on different networks (outside of LAN)
- Path selection between different source and destination
- [[network-devices#router|Router]] Works on this level

> [!FAQ] DATA (broken into Segments)
> SourceIP
> DestinationIP
>
> TCP
> Source-Port = 80
> Destination-Port = 7268
>
> HTTP Request
> GET https://getdata?page=1
## DataLink

Responsible for delivery within the same subnet
> [!FAQ] DATA (broken into Segments)
>
> Source-Mac (can be of a router's if out of subnet)
> Destination-Mac (can be of a router's if out of subnet)
>
> SourceIP
> DestinationIP
>
> TCP
> Source-Port = 80
> Destination-Port = 7268
>
> HTTP Request
> GET https://getdata?page=1
>
> (if Ethernet) Ethernet Trailer
### common-protocols
- Ethernet
- point to point protocol
## Physical


[[mac-addresses|mac-addresses]]

[[PDUs]]


