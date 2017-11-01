
[chroot](https://en.wikipedia.org/wiki/Chroot)
---
* An operation that changes the apparent  root directory for the current running process and its children.  


* let's get alpine filesystem
```
vagrant@vagrant:~$ docker run -it --name alpine-fs alpine touch /ALPINE
vagrant@vagrant:~$ docker export alpine-fs  -o alpine.tar
vagrant@vagrant:~$ mkdir /home/vagrant/alpine
vagrant@vagrant:~$ sudo tar xvf alpine.tar -C /home/vagrant/alpine
```

* let's chroot to this file system

```
sudo go run /vagrant/01-namespace/04-chroot/main.go run /bin/sh
```

[ns-in-docker](../05-ns-docker/README.md)