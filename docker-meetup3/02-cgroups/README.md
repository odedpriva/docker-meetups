Control Groups (cgroups)
---
provide a mechanism for easily managing and monitoring system resources, 
by partitioning things like cpu time, system memory, disk and network bandwidth, into groups, then assigning tasks to those groups

* cpu time
* system memory
* disk bandwidth
* network bandwidth
* monitoring

![image](https://sysadmincasts.com/static/extra/14-cgroups-overview.gif)


As an example of a scenario that can benefit from multiple hierarchies, consider a large
university server with various users - students, professors, system
tasks etc. The resource planning for this server could be along the
following lines:

       CPU :          "Top cpuset"
                       /       \
               CPUSet1         CPUSet2
                  |               |
               (Professors)    (Students)

               In addition (system tasks) are attached to topcpuset (so
               that they can run anywhere) with a limit of 20%

       Memory : Professors (50%), Students (30%), system (20%)

       Disk : Professors (50%), Students (30%), system (20%)

       Network : WWW browsing (20%), Network File System (60%), others (20%)
                               / \
               Professors (15%)  students (5%)


[let's start](00-setup/README.md)