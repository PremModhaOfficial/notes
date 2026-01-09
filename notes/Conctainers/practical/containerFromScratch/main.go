package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// docker 		  run img cmd params
// go run main.go run     cmd params
func main() {
	switch os.Args[1] {
	case "run":
		run()

	case "child":
		child()

	default:
		panic("unsupported command")
	}
}

func child() {
	fmt.Printf("Running cmd: %s with PID:%d\n", os.Args[2:], os.Getpid())

	syscall.Sethostname([]byte("prem's_container"))

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	must(cmd.Run())
}

func run() {
	fmt.Printf("Running cmd: %s with PID:%d\n", os.Args[2:], os.Getpid())

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// the syscall to set the namespace
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}

	must(cmd.Run())
}

func must(e error) {
	if e != nil {
		panic(e)
	}
}
