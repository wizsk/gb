package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/wizsk/gb/aes"
	"golang.org/x/term" // Import the term package for secure password input
)

// this will contain the file name
var tempFile string

func main() {
	if len(os.Args) != 4 {
		log.Fatal("not enoug args")
	}

	go shutdown()
	pass, err := getPassword()
	if err != nil {
		log.Fatal(err)
	}

	if os.Args[1] == "e" {
		err = aes.EncryptFile(os.Args[2], os.Args[3], pass)
	} else if os.Args[1] == "d" {
		err = aes.DecryptFile(os.Args[2], os.Args[3], pass)
	}
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(os.Args[2], "->", os.Args[3])
	}
}

func shutdown() {
	sighalChannel := make(chan os.Signal, 1)
	signal.Notify(sighalChannel, os.Interrupt, syscall.SIGTERM)

	// exiting gracefully
	<-sighalChannel
	if tempFile != "" {
		fmt.Println("\ncleaning temp files")
		err := os.Remove(tempFile)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Printf("\n%q exited\n", os.Args[0])
	os.Exit(0)
}

func getPassword() (string, error) {
	fmt.Print("Enter the password: ")
	password, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}
	fmt.Println() // Print a newline after the password input
	return string(password), nil
}
