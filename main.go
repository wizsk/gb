package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// this global var is used for cleaing up files incase of file wasn't cleaned
var tempFile string

func main() {
	// fmt.Println(os.UserConfigDir())
	// fmt.Println(os.UserCacheDir())
	// os.Exit(0)
	closeSig := make(chan struct{})
	go shutdown(closeSig)

	fl := os.Args[1]
	key := "hi man"

	err := createNewFile("tmp", fl, "nvim", key)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = openEncrypted("tmp", fl, "nvim", key)
	if err != nil {
		fmt.Println(err)
		return
	}

	// sending exit
	closeSig <- struct{}{}
	<-closeSig // it will wait shutdown func to complete
}

// func shutdown(closeSig <-chan struct{}) {
func shutdown(closeSig chan struct{}) {
	sighalChannel := make(chan os.Signal, 1)
	signal.Notify(sighalChannel, os.Interrupt, syscall.SIGTERM)

	// exiting gracefully
	select {
	case <-sighalChannel:
		break
	case <-closeSig:
		break
	}

	if tempFile != "" {
		fmt.Println("\ncleaning temp files")
		err := os.Remove(tempFile)
		if err != nil {
			fmt.Printf("while removing file '%s'\nerr: '%v'\n", tempFile, err)
		}
	}

	fmt.Printf("\n%q exited\n", os.Args[0])
	os.Exit(0)
	closeSig <- struct{}{}
}
