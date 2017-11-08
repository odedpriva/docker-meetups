### Namespaces
feature of the Linux kernel that isolate and virtualize system resources of a collection of processes.


#### avialable namespace 
As of kernel version 4.10 (Released 19 February, 2017), there are 7 kinds of namespaces.

Namespace | Clone Flag | Kernel | Year
----------|------------|--------|-----
MNT | CLONE_NEWNS | 2.4.19 | 2002
UTS | CLONE_NEWUTS | 2.6.19 | 2006
PID | CLONE_NEWPID | 2.6.24 | 2008
NET | CLONE_NEWNET | 2.6.29* | 2009
IPC | CLONE_NEWIPC | 2.6.30* | 2009
USER | CLONE_NEWUSER | 3.8* | 2013
CGROUP | CLONE_NEWCGROUP | 4.6* | 2016


![image](https://uploads.toptal.io/blog/image/674/toptal-blog-image-1416487554032.png)
![image](https://uploads.toptal.io/blog/image/675/toptal-blog-image-1416487605202.png)
![image](https://uploads.toptal.io/blog/image/677/toptal-blog-image-1416545619045.png)

[let's start](00-setup/README.md)

### external links

[introducing-namespaces](http://windsock.io/introducing-namespaces/)



the resources I used for this session: 
* http://windsock.io/introducing-namespaces/
* https://www.youtube.com/watch?v=MHv6cWjvQjM
* https://medium.com/@teddyking/namespaces-in-go-basics-e3f0fc1ff69a

---
 few notes : 
 * [fork linux](https://www.systutorials.com/docs/linux/man/2-fork/)
 * [linux namespaces](http://man7.org/linux/man-pages/man7/namespaces.7.html)
 * [linux chroot](http://man7.org/linux/man-pages/man2/chroot.2.html)
 * [/proc/[pid]/exe](http://man7.org/linux/man-pages/man5/proc.5.html) - Under Linux 2.2 and later, this file is a symbolic link containing the actual pathname of the executed command.  