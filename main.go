package main

import (
	"fmt"
	"os"

	"github.com/wizsk/gb/cmd"
)

const version = "0.1-dev"

func main() {
	tmpfile := ""
	root, err := cmd.RootCmd(&tmpfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	root.Version = "v" + version
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
