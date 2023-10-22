package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/wizsk/gb/aes"
	"golang.org/x/term" // for sequre pass input
)

// name mane file name
func openEncrypted(directory, name, editor, key string) error {
	if directory == "" || name == "" {
		return fmt.Errorf("openEncyped: directory name or file name is empty")
	}

	// create the tmp dir
	decFile, err := os.CreateTemp(directory, strings.TrimLeft(name, ".md.enc")+"-*.md")
	if err != nil {
		return err
	}
	decFile.Close()
	tempFile = decFile.Name()

	// delete the decrypted file here
	defer func() {
		if err := os.Remove(decFile.Name()); err == nil {
			// err will be caught by the shudown func
			tempFile = ""
		}
	}()

	// decrypt file
	err = aes.DecryptFile(filepath.Join(directory, name+".md.enc"), decFile.Name(), key)
	if err != nil {
		return err
	}

	err = openEditor(decFile.Name(), editor)
	if err != nil {
		return err
	}

	return aes.EncryptFile(decFile.Name(), filepath.Join(directory, name+".md.enc"), key)
}

func createNewFile(dir, fileName, editor, key string) error {
	if _, err :=os.Stat(filepath.Join(dir, fileName+".md.enc")); err == nil {
		return fmt.Errorf("createNewFile: %q already exists", fileName)
	}

	if dir == "" {
		return fmt.Errorf("createNewFile: directory name is empty")
	}

	if fileName == "" {
		fileName = time.Now().Format("03-04-05-PM-05-06-2006")
	}

	tf, err := os.CreateTemp(dir, fileName+"-*.md")
	if err != nil {
		return err
	}
	tf.Close()
	tempFile = tf.Name()
	// delete the decrypted file here
	defer func() {
		if err := os.Remove(tf.Name()); err == nil {
			// err will be caught by the shudown func
			tempFile = ""
		}
	}()

	if err = openEditor(tf.Name(), editor); err != nil {
		return err
	}

	stat, err := os.Stat(tf.Name())
	if err != nil {
		return err
	}

	if stat.Size() == 0 {
		return fmt.Errorf("createNewFile: noting was written to the file %q", fileName)
	}

	return aes.EncryptFile(tf.Name(), filepath.Join(dir, fileName+".md.enc"), key)
}

func openEditor(file, editor string) error {
	cmd := exec.Command(editor, file)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
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
