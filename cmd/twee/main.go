package main

import (
	"os"

	"github.com/kironono/twee/cmd/twee/commands"
)

func main() {
	if err := commands.Run(os.Args[1:]); err != nil {
		os.Exit(1)
	}
}
