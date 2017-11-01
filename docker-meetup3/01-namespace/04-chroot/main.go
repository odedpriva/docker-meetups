package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// go run main.go run <cmd> <args>
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

	must(cmd.Run())
}

func child() {
	fmt.Printf("Running child %v \n", os.Args[2:])
	name := "container(not really)"
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(syscall.Sethostname([]byte(name)))

	// chrooting to new file system
	must(syscall.Chroot("/home/vagrant/alpine"))
	must(os.Chdir("/"))
	// https://golang.org/pkg/syscall/#Mount
	// mounting proc.
	must(syscall.Mount("proc", "proc", "proc", 0, ""))

	must(cmd.Run())

	// unmounting proc
	must(syscall.Unmount("proc", 0))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
