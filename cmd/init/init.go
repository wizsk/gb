package init

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/wizsk/gb/aes"
	"github.com/wizsk/gb/config"
	"golang.org/x/term"
	"gopkg.in/yaml.v3"
)

const defatutNoteBook = "default"

// Init(nil) is the way to call
//
// the func is for testing
func Init(force bool, getPass func() (string, error)) error {
	root, err := config.RootDir()
	if err != nil {
		return err
	}

	// check if the file exists or not
	if _, err := os.Stat(root); err != nil && !os.IsNotExist(err) {
		return err
	}

	err = os.Mkdir(root, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		return err
	}

	configFile := filepath.Join(root, config.ConfigFileName)
	if _, err := os.Stat(configFile); err != nil && !os.IsNotExist(err) {
		return err
	} else if err == nil && !force {
		return fmt.Errorf("config %q already exists", configFile)
	}

	conf := config.DefaultConf()

	if getPass == nil {
		getPass = getPassword
	}

	pass, err := getPass()
	if err != nil {
		return err
	} else if pass == "" {
		return fmt.Errorf("password can not be empty")
	}
	conf.Key = aes.StringToHashHex(pass)

	confYml, err := yaml.Marshal(&conf)
	if err != nil {
		return err
	}

	if err = os.Mkdir(filepath.Join(root, defatutNoteBook), os.ModePerm); err != nil && !os.IsExist(err) {
		return err
	}

	fmt.Println(string(confYml))
	err = os.WriteFile(configFile, confYml, 0666)
	return nil
}

func getPassword() (string, error) {
	fmt.Println("Bear in mind that you will never be able to change the pass :) so give a strong one")
	fmt.Print("Enter the password: ")
	password, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}
	fmt.Println() // Print a newline after the password input
	return string(password), nil
}
