let's check how docker uses namespace

```bash

vagrant@vagrant:~$ docker run -d -e TEST=1000 alpine nc -l 9000

vagrant@vagrant:~$ sudo lsns | grep 9000
4026532197 mnt         1 22892 root            nc -l 9000
4026532198 uts         1 22892 root            nc -l 9000
4026532199 ipc         1 22892 root            nc -l 9000
4026532200 pid         1 22892 root            nc -l 9000
4026532202 net         1 22892 root            nc -l 9000
```

let's check what we can find out from /proc fs

we can check the command line of the process

```bash

sudo cat /proc/23290/environ  | tr '\0' '\n'
PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
HOSTNAME=2b5905ac633f
TEST=1000
HOME=/root

```

[cgroups](../../02-cgroups/README.md)