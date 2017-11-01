let's start by setting up the VM

```
vagrant up
```

let's start with checking that everything is working as expected.
remember to run this inside the vagrant machine you created

```  
vagrant ssh

vagrant@vagrant:~$ go run /vagrant/01-namespace/00-setup/main.go run echo hello
Running [echo hello]
hello
 
```

We got a running VM. now let's create out first [container](../01-uts/README.md)
 