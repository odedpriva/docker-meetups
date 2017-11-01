to make our 'container' start with a specific name, we need to create a 'child' process. 

now, let's start the 'container' and assign a hostname

```
sudo go run /vagrant/01-namespace/02-uts/main.go run /bin/bash
```

check the process list `ps aux`, what do you see? 

let's [give](../03-pid/README.md) it it's own process list  



[/proc/[pid]/exe](http://man7.org/linux/man-pages/man5/proc.5.html) - Under Linux 2.2 and later, this file is a symbolic link containing the actual pathname of the executed command. 