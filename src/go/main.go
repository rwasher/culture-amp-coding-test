package main

import (
	"fmt"
	"os"
)

func main() {
	Run(os.Args, os.Stdout, os.Stderr)
}

func Run(args []string, stdout, stderr *os.File) {
	oStdout := os.Stdout
	oStderr := os.Stderr
	defer func() {
		os.Stdout = oStdout
		os.Stderr = oStderr
	}()

	os.Stdout = stdout
	os.Stderr = stderr

	// Simplistic argument handling for starters
	if len(args) < 3 {
		fmt.Printf("Usage: %s <file path> <analysis mode>\n", os.Args[0])
		os.Exit(1)
	}

	file := args[1]
	mode := args[2]

	fmt.Printf("Run '%s' analysis on '%s'\n", mode, file)

	// Your implementation starts here
}
