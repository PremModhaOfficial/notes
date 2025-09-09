---
id: zmq
aliases:
  - ZMQ
tags: []
created: 2025-09-05 16:46:00
modified: 2025-09-09 17:20:46
---

# ZMQ

<!--toc:start-->
- [ZMQ](#zmq)
  - [chapter 4 Reliability](#chapter-4-reliability)
    - [What is Reliability](#what-is-reliability)
<!--toc:end-->

- the zmq is the built in way of managing the network comunication through the sockets
- the zmq will take kare of the things like quing and droping etc through its abstractions like the 

[[zmq-advanced-request-responce-patterns#important|zmq-advanced-request-responce-patterns]]


## chapter 4 Reliability
### What is Reliability
- the reliability is only understood in terms of failures,
- so if we can handle some set of well defined and well documented failures then we are reliable for thoes particular set of failures


The [[Lazy Pirate pattern]]: reliable request-reply from the client side
The Simple Pirate pattern: reliable request-reply using load balancing
The Paranoid Pirate pattern: reliable request-reply with heartbeating
The Majordomo pattern: service-oriented reliable queuing
The Titanic pattern: disk-based/disconnected reliable queuing
The Binary Star pattern: primary-backup server failover
The Freelance pattern: brokerless reliable request-reply


