
let's get alpine filesystem
```
docker run -it --name alpine-fs alpine touch /ALPINE
docker export alpine-fs  -o alpine.tar
mkdir /home/vagrant/alpine
sudo tar xvf alpine.tar -C /home/vagrant/alpine
```

let's chroot to this file system

```
sudo go run /vagrant/01-namespace/04-chroot/main.go run /bin/sh
```