For security reasons, it’s essential that the hostname has been set before the namespaced /bin/sh process starts running. 
After all, we don’t want programs running inside ns-process to be able to discover the Host’s hostname.

```go
cmd := exec.Command("/bin/echo", "Process already running")
cmd.SysProcAttr = &syscall.SysProcAttr{
 Cloneflags: syscall.CLONE_NEWUTS,
}
cmd.Run()
```

There’s no hook or anything here that allows us to run code after the namespace creation but before the process starts.


the naive way to achieve it is:  

```go
func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		fmt.Printf("support only run as a sub command, got %v\n", os.Args[1])
		os.Exit(0)
	}

}

func run() {
	fmt.Printf("Running %v \n", os.Args[2:])

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}

	cmd.Run()
}

func child() {
	fmt.Printf("Running child %v \n", os.Args[2:])
	name := "container(not really)"
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

    // for example, setting up hostname.
	syscall.Sethostname([]byte(name))

	cmd.Run()

}

```
to use the above code we could have run something like this:
```bash
naive run "sleep 1"
```

docker has implemented the same mechanism in a package called [reexec](https://github.com/moby/moby/blob/master/pkg/reexec/reexec.go)

let's use Docker's reexec package. 

```bash
vagrant@vagrant:~$ go run /vagrant/01-namespace/04-reexec/main.go

>> namespace setup code goes here <<

-[ns-process]- #

```

check the mount (`mount`) list or the process list (`ps aux`)  ? what do you see ? 

 