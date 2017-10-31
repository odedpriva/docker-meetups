

 so, now let's see why do we need the mount namespace 


```bash

vagrant@vagrant:~/alpine$ sudo go run /vagrant/01-namespace/05-mount/main.go run /bin/sh
/ # mount
proc on /proc type proc (rw,relatime)
mymount on /tmp type tmpfs (rw,relatime)

#running `docker run -it alpine mount` will not show the mounted fs for the process above.  
```

but, let's check them using /proc

```
sudo lsns
sudo cat /proc/23799/mounts
..

```

one more /proc 

```
container(not really): export TEST=luminate
container(not really): sleep 1000


vagrant@vagrant:~$ sudo cat /proc/$(pidof sleep)/environ | tr '\0' '\n' | grep TEST
TEST=luminate

```