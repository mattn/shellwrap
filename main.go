package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"unsafe"

	"github.com/mattn/go-shellwords"
)

func main() {
	us := syscall.GetCommandLine()
	p := (*[0xffff]uint16)(unsafe.Pointer(us))
	s := syscall.UTF16ToString(p[:])
	parser := shellwords.NewParser()
	parser.ParseEnv = true
	parser.ParseBacktick = true
	args, err := parser.Parse(s)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	args = args[1:]

	useshell := false
	if len(args) > 0 && args[0] == "-s" {
		useshell = true
		args = args[1:]
	}

	if len(args) == 0 {
		os.Exit(1)
	}

	var cmdargs []string
	if useshell {
		cmdargs = []string{"cmd", "/c"}
		cmdargs = append(cmdargs, args...)
	} else {
		cmdargs = args
	}
	cmd := exec.Command(cmdargs[0], cmdargs[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
