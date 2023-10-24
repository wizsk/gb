package main

import (
	"fmt"
	"os"

	"github.com/wizsk/gb/cmd"
	"github.com/wizsk/gb/config"
)

const version = "0.1-dev"
const debug = true

func main() {
	if debug {
		os.Setenv(config.RootDirEnvName, "./tmp")
	}

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
