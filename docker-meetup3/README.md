setting the environment
---
if you want to follow along with the examples, you need 
* [vagrant](https://www.vagrantup.com/)

```bash
docker-meetup3 git:(master) ✗ vagrant up
Bringing machine 'default' up with 'virtualbox' provider...
==> default: Importing base box 'bento/ubuntu-17.04'...
..
..

```


Container from scratch
---

![The future of Linux Containers](https://www.youtube.com/watch?v=wW9CAH9nSLs)

Often thought of as cheap VMs, containers are just isolated groups of processes running on a single host. 

* Operating-system-level virtualization, also known as containerization 
* Operating system feature in which the kernel allows the existence of multiple isolated user-space instances
* Programs running inside a container can only see the container's contents and devices assigned to the container.

Linux Containers: 
---
uses discrete kernel features like cgroups, namespaces, SELinux, and more.

* [namespace](https://en.wikipedia.org/wiki/Linux_namespaces)  
* [cgroup](https://en.wikipedia.org/wiki/Cgroups)


OCI
---
* OCI ( open container initiative ) - Established in June 2015 by Docker and other leaders in the container industry
* OCI currently contains two specifications: the Runtime [Specification](http://www.github.com/opencontainers/runtime-spec) and the Image [Specification](http://www.github.com/opencontainers/image-spec)

![OCI](https://i0.wp.com/blog.docker.com/wp-content/uploads/243938a0-856b-4a7f-90ca-2452a69a385c-1.jpg?w=1019&ssl=1)

OCI implementations: 
* [runc](https://github.com/opencontainers/runc)
* [railcar](https://github.com/oracle/railcar) - Rust implementation of the Open Containers Initiative oci-runtime

let's start with [namespace](01-namespace/README.md)



note: external links
---  
* [oci](https://www.opencontainers.org/)  
* [docker and oci](https://blog.docker.com/2017/07/demystifying-open-container-initiative-oci-specifications/)  





