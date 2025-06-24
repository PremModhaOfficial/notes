---
id: ps
aliases:
  - prosses running in current shell
  - prosses running in current shell identifiable by tty's 'pts/1' meaning the psudo terminal
tags: []
created: 2025-07-01 14:58:46
modified: 2025-07-02 11:57:02
---


# prosses running in current shell
- identifiable by tty's `pts/1` meaning the psudo terminal
```sh 
$ ps
    PID TTY          TIME CMD
 118349 pts/1    00:00:00 fish
 119476 pts/1    00:00:00 ps
```

# Most widely used
- ps -e or ps -A ((IDENTICAL))

## sorted b pid all the prosseses
```sh
 ps -e | head
    PID TTY          TIME CMD
      1 ?        00:00:02 systemd
      2 ?        00:00:00 kthreadd
      3 ?        00:00:00 pool_workqueue_release
      4 ?        00:00:00 kworker/R-rcu_gp
      5 ?        00:00:00 kworker/R-sync_wq
      6 ?        00:00:00 kworker/R-slub_flushwq
      7 ?        00:00:00 kworker/R-netns
      9 ?        00:00:00 kworker/0:0H-events_highpri
     12 ?        00:00:00 kworker/R-mm_percpu_wq

```

## F for full format
```sh
$ ps -ef | head
UID          PID    PPID  C STIME TTY          TIME CMD
root           1       0  0 10:31 ?        00:00:02 /sbin/init splash
root           2       0  0 10:31 ?        00:00:00 [kthreadd]
root           3       2  0 10:31 ?        00:00:00 [pool_workqueue_release]
root           4       2  0 10:31 ?        00:00:00 [kworker/R-rcu_gp]
root           5       2  0 10:31 ?        00:00:00 [kworker/R-sync_wq]
root           6       2  0 10:31 ?        00:00:00 [kworker/R-slub_flushwq]
root           7       2  0 10:31 ?        00:00:00 [kworker/R-netns]
root           9       2  0 10:31 ?        00:00:00 [kworker/0:0H-events_highpri]
root          12       2  0 10:31 ?        00:00:00 [kworker/R-mm_percpu_wq]

```


# BSD format that is giving more information than -A and -e
```sh
 ps aux | head
USER         PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
root           1  0.0  0.0  23940 14452 ?        Ss   10:31   0:02 /sbin/init splash
root           2  0.0  0.0      0     0 ?        S    10:31   0:00 [kthreadd]
root           3  0.0  0.0      0     0 ?        S    10:31   0:00 [pool_workqueue_release]
root           4  0.0  0.0      0     0 ?        I<   10:31   0:00 [kworker/R-rcu_gp]
root           5  0.0  0.0      0     0 ?        I<   10:31   0:00 [kworker/R-sync_wq]
root           6  0.0  0.0      0     0 ?        I<   10:31   0:00 [kworker/R-slub_flushwq]
root           7  0.0  0.0      0     0 ?        I<   10:31   0:00 [kworker/R-netns]
root           9  0.0  0.0      0     0 ?        I<   10:31   0:00 [kworker/0:0H-events_highpri]
root          12  0.0  0.0      0     0 ?        I<   10:31   0:00 [kworker/R-mm_percpu_wq]

- [ ] ```
