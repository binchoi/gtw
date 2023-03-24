package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("go", "test", "-v")

	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", stdoutStderr)
}
