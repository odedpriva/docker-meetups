let's start by setting up the VM

```bash
docker-meetup3 git:(master) ✗ vagrant up
....
```

let's start with checking that everything is working as expected.

```bash
docker-meetup3 git:(master) ✗ vagrant ssh
```

Creating namespace is super easy, just a single syscall with one argument, [unshare](https://linux.die.net/man/2/unshare)

The unshare command allows you to run a program with some namespaces ‘unshared’ from its parent. Essentially what this means is that unshare will run whatever program you pass it in a new set of namespaces.

```bash
vagrant@vagrant:~$ unshare -h

Usage:
 unshare [options] <program> [<argument>...]

Run a program with some namespaces unshared from the parent.

Options:
 -m, --mount[=<file>]      unshare mounts namespace
 -u, --uts[=<file>]        unshare UTS namespace (hostname etc)
 -i, --ipc[=<file>]        unshare System V IPC namespace
 -n, --net[=<file>]        unshare network namespace
 -p, --pid[=<file>]        unshare pid namespace
 -U, --user[=<file>]       unshare user namespace
 -C, --cgroup[=<file>]     unshare cgroup namespace
 -f, --fork                fork before launching <program>
     --mount-proc[=<dir>]  mount proc filesystem first (implies --mount)
 -r, --map-root-user       map current user to root (implies --user)
     --propagation slave|shared|private|unchanged
                           modify mount propagation in mount namespace
 -s, --setgroups allow|deny  control the setgroups syscall in user namespaces

 -h, --help     display this help and exit
 -V, --version  output version information and exit

```

```bash

vagrant@vagrant:~$ sudo su - # become root user sudo /etc/profile, .profile and .bashrc are executed #
root@vagrant:~ unshare -u bash # create a shell in new UTS namespace
root@vagrant:~ hostname oded # set hostname
root@vagrant:~ hostname # confirm new hostname
oded
root@vagrant:~# exit # exit new UTS namespace
exit
root@vagrant:~# hostname # confirm original hostname unchanged
vagrant

```

```bash
# Establish a PID namespace.
vagrant@vagrant:~$ sudo unshare --fork --pid --mount-proc bash
root@vagrant:~ ps aux
USER       PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
root         1  0.0  0.3  23704  3744 pts/0    S    19:55   0:00 bash
root        10  0.0  0.3  40056  3148 pts/0    R+   19:55   0:00 ps aux

# list namepspaces
root@vagrant:~# sudo lsns
        NS TYPE   NPROCS   PID USER            COMMAND
4026531835 cgroup    117     1 root            /sbin/init
...
4026532195 mnt         2 20862 root            unshare --fork --pid --mount-proc bash
4026532196 pid         1 20863 root            bash

# another tool that enable running process in the context of the name space is `nsenter`
root@vagrant:~# nsenter --target 20863 --mount --uts --ipc --net --pid
root@vagrant:/# ps uax
USER       PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
root         1  0.0  0.3  23704  3736 pts/0    S+   20:15   0:00 bash
root        11  0.0  0.3  23712  3496 pts/1    S    20:15   0:00 -bash
root        22  0.0  0.3  40056  3128 pts/1    R+   20:15   0:00 ps uax
```


```bash
# Establish a PID namespace - process is getting PID #1
vagrant@vagrant:~$ sudo unshare --fork --pid --mount-proc readlink /proc/self
1
```

[now](../01-uts/README.md), we want to make it change the hostname befor we 'enter' the 'container'

---
[proc](https://en.wikipedia.org/wiki/Procfs) - a special filesystem in Unix-like operating systems that presents information about processes and other system information in a hierarchical file-like structure


 