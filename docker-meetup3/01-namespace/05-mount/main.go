/*

Mount namespaces provide isolation of the list of mount points seen
by the processes in each namespace instance.  Thus, the processes in
each of the mount namespace instances will see distinct single-
directory hierarchies.

*/

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
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
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

	must(syscall.Chroot("/home/vagrant/alpine"))
	must(os.Chdir("/"))
	// let's add a nice prompt
	must(os.Setenv("PS1", name+": "))
	must(syscall.Mount("proc", "proc", "proc", 0, ""))

	// tmpfs - It is intended to appear as a mounted file system, but stored in volatile memory instead of a persistent storage device
	must(syscall.Mount("mymount", "/tmp", "tmpfs", 0, ""))

	must(cmd.Run())

	must(syscall.Unmount("proc", 0))
	must(syscall.Unmount("/tmp", 0))

}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
