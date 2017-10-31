
let's seperate the process list by using `CLONE_NEWPID` system call
A PID namespace isolates process ID (PID) numbers, thereby allowing processes running on the same system to have the same PID.

```
sudo go run /vagrant/01-namespace/03-pid/main.go run /bin/bash

ls -la /proc 
# now, what do you see ?

```
ps aux is lookikng inside /proc direcotry

so, if we want the 'container' to have it's own process list, we need to create a copy of /proc

for that we use chroot 