[UTS](http://windsock.io/uts-namespace/)
---
The UTS namespace is used to isolate two specific elements of the system that relate to the uname system call
for our purpose, it will change the hostname without affecting the host's hostname

let's add a clone flag that allow the change of the hostname inside the namespace.
```
go run /vagrant/01-namespace/01-uts/main.go run /bin/bash
Running [/bin/bash]
panic: fork/exec /bin/bash: operation not permitted

goroutine 1 [running]:
.
.

```

first thing we see, is that we need root permission to use namespeces capabilities.

let's add sudo. 
```
sudo go run /vagrant/01-namespace/01-uts/main.go run /bin/bash
Running [/bin/bash]
root@vagrant:~# hostname luminate
root@vagrant:~# hostname
luminate
```

[now](../02-uts/README.md), we want to make it change the hostname befor we 'enter' the 'container'