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
	fmt.Printf("Running main %v \n", os.Args[2:])

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS,
	}

	must(cmd.Run())
}

func child() {
	fmt.Printf("Running chile %v \n", os.Args[2:])
	name := "container(not really)"
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// setting the host name insid the 'container'
	must(syscall.Sethostname([]byte(name)))

	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
