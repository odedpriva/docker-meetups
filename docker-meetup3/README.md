Container from scratch
---
* Operating-system-level virtualization, also known as containerization 
* Operating system feature in which the kernel allows the existence of multiple isolated user-space instances
* Programs running inside a container can only see the container's contents and devices assigned to the container.

OCI
---
* OCI ( open container initiative ) - Established in June 2015 by Docker and other leaders in the container industry
* OCI currently contains two specifications: the Runtime [Specification](http://www.github.com/opencontainers/runtime-spec) and the Image [Specification](http://www.github.com/opencontainers/image-spec)

![OCI](https://i0.wp.com/blog.docker.com/wp-content/uploads/243938a0-856b-4a7f-90ca-2452a69a385c-1.jpg?w=1019&ssl=1)

Linux Containers: 
---
uses discrete kernel features like cgroups, namespaces, SELinux, and more.

* [namespace](namespace.md)  
* [cgroups](cgroups.md)  




note: external links
---
* [namespace](https://en.wikipedia.org/wiki/Linux_namespaces)  
* [cgroup](https://en.wikipedia.org/wiki/Cgroups)  
* [oci](https://www.opencontainers.org/)  
* [docker and oci](https://blog.docker.com/2017/07/demystifying-open-container-initiative-oci-specifications/)  





