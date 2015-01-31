package main

import (
	"fmt"
)

func main() {
	cmdPath := "/bin/echo"
	cmdArgs := []string{"Hello", "World"}
	cmdEnv := []string{}

	fmt.Printf("path: %v\nargs: %v\nenv: %v\n", cmdPath, cmdArgs, cmdEnv)
}
