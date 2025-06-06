---
id: 1749186058-osi
aliases:
  - osi
tags: []
---

# osi

It is a refrence model for machines to comunicate via Electronic means
some machine may not exactly mimic the model
all of these layers talk to the layers adjacent to them

all layers wrap and unwrap the data breaks into packets



# modern TCP/IP
## Application
Http Protocol lives here

> [!FAQ] DATA
> HTTP Request
> GET https://getdata?page=1

## Presentation
## Session
## Transport

Port Numbers To track sessions
> [!FAQ] DATA
> TCP
> Source-Port = 80
> Destination-Port = 7268
> HTTP Request
> GET https://getdata?page=1

## Network
IP addresses
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

responsible for delivery within the same subnet
> [!FAQ] DATA (broken into Segments)
>
> Source-Mac       (can be of a router's if out of subnet)
> Destination-Mac  (can be of a router's if out of subnet)
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
> (if ethernet) Ethernet Trailer
### commont-protocols
- ethernet
- point to point protocol
## Physical


[[mac-addresses|mac-addresses]]
