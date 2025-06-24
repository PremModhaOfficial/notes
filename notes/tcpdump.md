---
id: tcpdump
aliases:
  - tcpdump
tags: []
created: 2025-06-27 10:49:00
modified: 2025-06-27 12:36:06
---

# tcpdump

help calture network traffic on your devices
- has a very simple pythonlike [[#filtering]]


## FLAGS
- c -- only capture this many packets
- i -- interface to capture from
- D -- list capturable devises
- S Sequence numbers
- s snapshot length how much of a packet to capture
- w -- write to file 
- r -- read from file / open file
- t - no time
- tt - time in unix epoc
- ttt - time diff b/w packets
- tttt - human readable
- ttttt - time diff b/w first execution 
- v/vv/vvv -- verbosity (2) will decode smps (mails)



## Parce it
`timestamp` `from_net_device` `src address` > `dest address`: `protocol (udp/tcp)`, `packet length`

## filtering
- host `IP` (both src and dst)
- dst `IP`
- src `IP` / port `ssh/22`
- `UDP`
- `TCP`
- net `CIDR` (subnet) [[CIDR]]
- port
- portrange `a-b` (from port a to b)

### oparators
- &&
- !
- ||
- `(`expr`)` [!note] can reqire you to wrap in `"`(expr)`"`

[[netstat]]



