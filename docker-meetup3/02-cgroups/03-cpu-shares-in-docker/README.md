let's see implementation of cgroups in Docker

run a container and a long running process
```
docker run -it --cpu-shares=512 ubuntu:17.04 /bin/bash

root@1d4ca31dd3dc:/# sleep 2000
```


let's see the cgroup in docker machine
```
# enter docker-for-mac / docker-for-windows 
docker run --rm -it --privileged --pid=host walkerlee/nsenter -t 1 -m -u -i -n sh
pidof sleep
2193


cat /sys/fs/cgroup/cpu/docker-ce/docker/1d4ca31dd3dc41f795c6e88ecd68808ede7e0244190eb54cbfae41a94a3c1251/tasks
2193
```