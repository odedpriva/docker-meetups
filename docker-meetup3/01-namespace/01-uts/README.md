[UTS](http://windsock.io/uts-namespace/)
---

* The UTS namespace is used to isolate two specific elements of the system that relate to the uname system call
* For our purpose, it will change the hostname without affecting the host's hostname

let's add a clone flag that allow the change of the hostname inside the namespace.
```
vagrant@vagrant:~$ go run /vagrant/01-namespace/01-setup/main.go
Error running the /bin/sh command - fork/exec /bin/sh: operation not permitted
exit status 1

```

first thing we see, is that we need root permission to use namespeces capabilities.

let's add sudo. 
```bash
vagrant@vagrant:~$ sudo go run /vagrant/01-namespace/01-uts/main.go
-[ns-process]- hostname
vagrant
-[ns-process]- hostname oded
-[ns-process]- hostname
oded
-[ns-process]- exit
```

[next->](../02-uts/README.md)