
[chroot](https://en.wikipedia.org/wiki/Chroot)
---
* An operation that changes the apparent root directory for the current running process and its children.  


* let's get alpine filesystem
```
vagrant@vagrant:~$ docker run -it --name alpine-fs alpine touch /ALPINE
vagrant@vagrant:~$ docker export alpine-fs  -o alpine.tar
vagrant@vagrant:~$ mkdir /home/vagrant/alpine
vagrant@vagrant:~$ sudo tar xvf alpine.tar -C /home/vagrant/alpine
```

* let's chroot to this file system

```bash
vagrant@vagrant:~$ sudo go run /vagrant/01-namespace/05-chroot/main.go run /bin/sh
Running [/bin/sh]
Running child [/bin/sh]
/ ps aux
PID   USER     TIME   COMMAND
    1 root       0:00 /proc/self/exe child /bin/sh
    4 root       0:00 /bin/sh
    5 root       0:00 ps aux
/ mount
proc on /proc type proc (rw,relatime)
proc on /proc type proc (rw,relatime)
/ exit
```

[ns-in-docker](../05-ns-docker/README.md)