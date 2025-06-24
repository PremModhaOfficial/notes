---
id: systemd
aliases:
  - Systemd
tags: []
created: 2025-06-27 14:58:30
modified: 2025-06-27 16:34:02f
---

# Systemd

## UNITs
- managable resoureses of the system
- encludes but not limited to..  ((managed by the systen d))
  * mounts
  * automaount
  * timers
  * services

### Services ((systemctl command))
#### States of Service

```sh
systemctl status servise_name
```
- `disabled`
- `enabled`
- preset: `enabled`/`disabled`
#### Subcommands
```sh
systemctl start servise_name
systemctl restart servise_name
systemctl disabled servise_name
systemctl enabled servise_name
```
