---
id: awk
aliases:
  - os
  - awk
tags:
  - topic/os
  - concept/os
  - source/youtube/distrotube
created: 2025-06-27 14:57:53
modified: 2025-06-30 17:55:31
status:
  - #draft
type:
  - #atomic-note
---

# awk


### What can it do??
management of structured strings
- collumns
- logs
- csv type programeticaly saparate

provides declarative cli that helps manipulate the structure of the text

### How it works

- provide a saparator by -F (feild saparator) (default " ") it will make colummns out of strings using the saparator
- now at last decalre action ex: `"{print $1}"`

- case: `"{print $n}"` print nth collumn with saparator
- case: `"{print $NF}"` print last collimn (helpfull when the )
- case: `awk -F "/" '/^\// {print $NF}' /etc/shells` this prints the last saparated walues (shells) that are atvariable number of collumns


- `~` reg match
- `=` exact match
### Proppesr Scripting laguage
- case: `awk -F "/" '/^\// {print $NF}' /etc/shells` this prints the last saparated walues (shells) that are atvariable number of collumns

```bash
◎ awk 'BEGIN{ FS=" "; OFS="-->"} {print $1,length()}' /etc/shells
#-->33
/bin/sh-->7
/usr/bin/sh-->11
/bin/bash-->9
/usr/bin/bash-->13
/bin/rbash-->10
/usr/bin/rbash-->14
/usr/bin/dash-->13
/usr/bin/fish-->13
/bin/zsh-->8
/usr/bin/zsh-->12
/usr/bin/tmux-->13

```

```bash
$ ps -ef | awk '{ if($NF ~ /\/usr\/bin/) print $0 }'

prem-mo+   54458   54441  0 13:27 ?        00:00:00 /usr/bin/pipewire
prem-mo+   54463   54441  0 13:27 ?        00:00:03 /usr/bin/wireplumber
prem-mo+   54464   54441  0 13:27 ?        00:00:01 /usr/bin/pipewire-pulse
prem-mo+   54741   54441  2 13:27 ?        00:03:15 /usr/bin/gnome-shell
prem-mo+   55287   54441  0 13:28 ?        00:00:00 /snap/snapd-desktop-integration/253/usr/bin/snapd-desktop-integration
prem-mo+   55348   55287  0 13:28 ?        00:00:00 /snap/snapd-desktop-integration/253/usr/bin/snapd-desktop-integration
prem-mo+   56510   55956  0 13:28 pts/2    00:00:00 /usr/bin/fish
prem-mo+   56666   55956  0 13:28 pts/4    00:00:00 fish -c cat '/home/prem-modha/.local/share/tmux/resurrect/restore/pane_contents//pane-tmux:1.1'; exec /usr/bin/fish
prem-mo+   58515   54694  0 13:28 ?        00:00:00 /usr/bin/update-notifier

```


[[bc_test]]
