---
id: nats-jetstream
aliases: []
tags: []
backlinks: []
created: 2026-01-06 15:48:03
links: []
modified: 2026-01-06 20:05:32
status: reviewed
---

## Jetstream

- Jetstream is a component added on top of the core for persistance of the messages

## whats the difference

| Before | After |
| -------------- | --------------- |
| at-most one delivery | at-most once, exactly-once and atleast-once  |
| no-persistance | messages-persist with diffrent [retention](nats-jetstream-retention) policies |


## Hands On
[Limit-based-jetstream](./NATS/practice/nats-jet-pract/limitBased/limitBased.md)

