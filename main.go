package main

import (
	"os"

	"learn-cli/cmd"
)

func main() {
	if err := cmd.ExecuteArgs(os.Args[1:], os.Stdout); err != nil {
		os.Exit(1)
	}
}
