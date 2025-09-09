---
id: nats-core
aliases: []
tags: []
backlinks: []
created: 2026-01-06 15:20:16
links: []
modified: 2026-01-06 18:15:14
status: reviewed
---


# Core nats

## Topics

*A subtring separeted by the dots is the Token*
*A Topic is basically a string made up of Tokens*
*This will form a hierchy of Topics*
*Core nats provides variouse messagint patterns on this topics*

### Subject-Matching
#### with "*"
- matches one token 
- "topic.a.b.*" will match "topic.a.b.anything" "topic.a.b.abcd" "topic.a.b.something" 

#### with ">"
- matches any number of proceeding tokens
- "topic.>" will match 
    * "topic.a.b.anything" "topic.a.b.abcd" "topic.a.b.something" 
    * "topic.X.b.anything" "topic.Y.b.abcd" "topic.Z.b.something" 
    * "topic.a.LOL" "topic.ABCD" "topic.a.ADDRS" 


## Messaging Patterns
[[nats-core-messaging-patterns]]


## HOW to name them??
- avoid the naming in such way that can cause the security problems (ie.)
    * "motadata.chats.prem"
        + any persont subscribed to the motadata.chats.* will see everyone's Names, msgs, and informations
    * "prem.chats.motadata"
        + this way the private things are at the left sides and all the public and general things are at the right side of the topics
