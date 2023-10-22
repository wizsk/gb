//go:build exclude

package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/wizsk/gb/aes"
	"github.com/wizsk/gb/config"
	"golang.org/x/term" // for sequre pass input
)

// name mane file name
func openEncrypted(c *config.Config, filename string) error {
	// create the tmp dir
	decFile, err := os.CreateTemp(c.RootDir, strings.TrimLeft(filename, ".md.enc")+"-*.md")
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
	err = aes.DecryptFile(c.FullEncFilePath(filename), decFile.Name(), c.Key)
	if err != nil {
		return err
	}

	err = openEditor(decFile.Name(), c.Editor)
	if err != nil {
		return err
	}

	return aes.EncryptFile(decFile.Name(), c.FullEncFilePath(filename), c.Key)
}

func createNewFile(c *config.Config, fileName string) error {
	if _, err := os.Stat(filepath.Join(c.RootDir, c.AddEncExt(fileName))); err == nil {
		return fmt.Errorf("createNewFile: %q already exists", fileName)
	}

	if fileName == "" {
		fileName = time.Now().Format("03-04-05-PM-05-06-2006")
	}

	tf, err := os.CreateTemp(c.RootDir, fileName+"-*.md")
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

	if err = openEditor(tf.Name(), c.Editor); err != nil {
		return err
	}

	stat, err := os.Stat(tf.Name())
	if err != nil {
		return err
	}

	if stat.Size() == 0 {
		return fmt.Errorf("createNewFile: noting was written to the file %q", fileName)
	}

	return aes.EncryptFile(tf.Name(), c.FullEncFilePath(fileName), c.Key)
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
