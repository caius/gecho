package main

import (
	"syscall"
)

func main() {
	cmdArgs := []string{"/bin/echo", "Hello", "World"}
	cmdEnv := []string{}

	err := syscall.Exec(cmdArgs[0], cmdArgs, cmdEnv)
	if err != nil {
		panic(err)
	}
}
