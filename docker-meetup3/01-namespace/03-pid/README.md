[PID](http://windsock.io/pid-namespace/)
---
A PID namespace isolates process ID (PID) numbers, thereby allowing processes running on the same system to have the same PID.

let's seperate the process list by using `CLONE_NEWPID` system call

```
sudo go run /vagrant/01-namespace/03-pid/main.go run /bin/bash

# let's check our process ID 

root@container(not really):~# echo $$
4

# and check the process list

root@container(not really):~# ps aux
USER       PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
root         1  0.0  0.7  57596  7268 ?        Ss   05:38   0:01 /sbin/init
root         2  0.0  0.0      0     0 ?        S    05:38   0:00 [kthreadd]
root         4  0.0  0.0      0     0 ?        S<   05:38   0:00 [kworker/0:0H]
root         5  0.0  0.0      0     0 ?        S    05:38   0:00 [kworker/u2:0]
root         6  0.0  0.0      0     0 ?        S    05:38   0:00 [ksoftirqd/0]
root         7  0.0  0.0      0     0 ?        S    05:38   0:00 [rcu_sched]
root         8  0.0  0.0      0     0 ?        S    05:38   0:00 [rcu_bh]
root         9  0.0  0.0      0     0 ?        S    05:38   0:00 [migration/0]
root        10  0.0  0.0      0     0 ?        S<   05:38   0:00 [lru-add-drain]
root        11  0.0  0.0      0     0 ?        S    05:38   0:00 [watchdog/0]
root        12  0.0  0.0      0     0 ?        S    05:38   0:00 [cpuhp/0]
root        13  0.0  0.0      0     0 ?        S    05:38   0:00 [kdevtmpfs]
root        14  0.0  0.0      0     0 ?        S<   05:38   0:00 [netns]
root        15  0.0  0.0      0     0 ?        S    05:38   0:00 [khungtaskd]
root        16  0.0  0.0      0     0 ?        S    05:38   0:00 [oom_reaper]
root        17  0.0  0.0      0     0 ?        S<   05:38   0:00 [writeback]
root        18  0.0  0.0      0     0 ?        S    05:38   0:00 [kcompactd0]
root        19  0.0  0.0      0     0 ?        SN   05:38   0:00 [ksmd]
root        20  0.0  0.0      0     0 ?        SN   05:38   0:00 [khugepaged]
root        21  0.0  0.0      0     0 ?        S<   05:38   0:00 [crypto]
root        22  0.0  0.0      0     0 ?        S<   05:3

.
.
.


```

why? ps aux is lookikng inside /proc direcotry

so, if we want the 'container' to only see the its own process, it needs to have its own copy of /proc, 

let's do that by using [chroot](../04-chroot/README.md)