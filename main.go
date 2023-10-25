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
		r := "./tmp"
		fmt.Printf("INFO: running in debug mode %q was set to %q\n", config.RootDirEnvName, r)
		os.Setenv(config.RootDirEnvName, r)
	}

	root, err := cmd.RootCmd()

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
