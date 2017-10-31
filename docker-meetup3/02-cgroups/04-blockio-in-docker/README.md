After version 1.10 docker added new features to manipulate IO speed in the container.
```

docker help run | grep -E 'bps|IO'
Usage:	docker run [OPTIONS] IMAGE [COMMAND] [ARG...]
      --blkio-weight uint16            Block IO (relative weight), between 10 and 1000, or 0 to disable (default 0)
      --blkio-weight-device list       Block IO weight (relative device weight) (default [])
      --device-read-bps list           Limit read rate (bytes per second) from a device (default [])
      --device-read-iops list          Limit read rate (IO per second) from a device (default [])
      --device-write-bps list          Limit write rate (bytes per second) to a device (default [])
      --device-write-iops list         Limit write rate (IO per second) to a device (default [])

```

run a container with write constraint and another one with no constraint
```
docker run -it --device-write-bps /dev/sda:50mb ubuntu:17.04 /bin/bash
root@e66cfc990bf2:/# time dd if=/dev/zero of=/tmp/file-1 bs=1M count=3000 oflag=direct status=progress

```

```
docker run -it ubuntu:17.04 /bin/bash
root@9c780bc83be0:/# time dd if=/dev/zero of=/tmp/file-1 bs=1M count=3000 oflag=direct status=progress

```
