
```
container(not really): sleep 1000

vagrant@vagrant:/sys/fs/cgroup/pids/vagrant$ cat cgroup.procs  | grep $(pidof sleep)
24344

vagrant@vagrant:/sys/fs/cgroup/pids/vagrant$ cat pids.max
20
```


let's test this rule by running a `fork bomb`

`:(){ :|: & };:` - DO NOT RUN THIS ON YOUR LOCAL MACHINE!!!!!!

```sh

root@container(not really):/# :() { : | : & }; :

# get the pid #
vagrant@vagrant:/sys/fs/cgroup/pids/vagrant$ sudo lsns 

# check # of processes in the pid
vagrant@vagrant:/sys/fs/cgroup/pids/vagrant$ pstree -H 24728
systemd-+-VBoxService-+-{automount}
        |             |-{control}
        |             |-{cpuhotplug}
        |             |-{memballoon}
        |             |-{timesync}
        |             |-{vminfo}
        |             `-{vmstats}
        |-accounts-daemon-+-{gdbus}
        |                 `-{gmain}
        |-acpid
        |-agetty
        |-atd
        |-cron
        |-dbus-daemon
        |-dhclient
        |-dockerd-+-docker-containe---9*[{docker-containe}]
        |         `-9*[{dockerd}]
        |-2*[iscsid]
        |-lvmetad
        |-lxcfs---2*[{lxcfs}]
        |-polkitd-+-{gdbus}
        |         `-{gmain}
        |-rpcbind
        |-rsyslogd-+-{in:imklog}
        |          |-{in:imuxsock}
        |          `-{rs:main Q:Reg}
        |-snapd---6*[{snapd}]
        |-sshd-+-sshd---sshd---bash---sudo---go-+-main-+-exe-+-47*[bash]
        |      |                                |      |     `-2*[{exe}]
        |      |                                |      `-2*[{main}]
        |      |                                `-4*[{go}]

```