[USER](http://man7.org/linux/man-pages/man7/user_namespaces.7.html)
---
* The User namespace provides isolation of UIDs and GIDs
* There can be multiple, distinct User namespaces in use on the same host at any given time
* Every Linux process runs in one of these User namespaces
* User namespaces allow for the UID of a process in User namespace 1 to be different to the UID for the same process in User namespace 2
* UID/GID mapping provides a mechanism for mapping IDs between two separate User namespaces

![image](https://cdn-images-1.medium.com/max/1600/1*lY9jQy-ZHnKF1fMEe0W9qQ.jpeg)

```
vagrant@vagrant:~$ go run /vagrant/01-namespace/03-user/main.go
-[ns-process]- # id
uid=0(root) gid=0(root) groups=0(root),65534(nogroup)
-[ns-process]- #

```

[next->](../04-reexec/README.md)