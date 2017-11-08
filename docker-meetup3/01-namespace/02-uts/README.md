now, let's start the 'container' and assign a hostname

```bash
vagrant@vagrant:~$ go run /vagrant/01-namespace/02-uts/main.go
-[ns-process]- id
uid=65534(nobody) gid=65534(nogroup) groups=65534(nogroup)
-[ns-process]- ip a
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
-[ns-process]- ps uax
..
..
..

-[ns-process]- mount
```

why we didn't needed to use sudo? 

* We’ve requested a new Mount namespace (CLONE_NEWNS) but are currently piggybacking off the host's mounts and rootfs
* We’ve requested a new PID namespace (CLONE_NEWPID) but haven't mounted a new /proc filesystem
* We’ve requested a new Network namespace (CLONE_NEWNET) but haven't setup any interfaces inside the namespace
* We’ve requested a new User namespace (CLONE_NEWUSER) but have failed to provide a UID/GID mapping

first, let's [fix](../03-user/README.md) UID and GID

---