package initGb

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/wizsk/gb/aes"
	"github.com/wizsk/gb/config"
	"github.com/wizsk/gb/core"
	"golang.org/x/term"
	"gopkg.in/yaml.v3"
)

func InitGb() *cobra.Command {
	i := &cobra.Command{
		Use:   "init",
		Short: "initialize gb",
	}

	var force bool
	i.Flags().BoolVarP(&force, "force", "f", false, "delete the old config and create a new one")

	i.RunE = func(cmd *cobra.Command, args []string) error {
		return initGB(force, getPassword)
	}

	return i
}

// if force == true then the previous confing will be rewritten
func initGB(force bool, getPass func() (string, error)) error {
	root, err := config.RootDir()
	if err != nil {
		return err
	}

	// check if the file exists or not
	if _, err := os.Stat(root); err != nil && !os.IsNotExist(err) {
		return err
	}

	if force {
		reader := bufio.NewReader(os.Stdin)

		for {
			fmt.Print("Do you really want to delete all your notes? [y/N] ")
			input, err := reader.ReadString('\n')
			if err != nil {
				return err
			}

			// Remove the newline character at the end of the input
			input = strings.TrimSuffix(input, "\n")
			input = strings.ToLower(input)
			fmt.Println()

			if input == "yes" || input == "y" {
				if err := os.RemoveAll(root); err != nil && !os.IsNotExist(err) {
					return err
				}
				break
			} else if input == "" || input == "no" || input == "n" {
				fmt.Println("No note was deleted")
				return nil
			}
		}
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

	pass, err := getPass()
	if err != nil {
		return err
	} else if pass == "" {
		return fmt.Errorf("password can not be empty")
	}

	keyFile := filepath.Join(root, config.KeyFileName)
	if err = os.WriteFile(keyFile, []byte(aes.StringToHashHex(pass)), 0666); err != nil {
		return err
	}
	fmt.Printf("Key written to %q\n", keyFile)

	if err = os.WriteFile(filepath.Join(root, ".gitignore"), []byte(config.KeyFileName+"\n"), 0666); err != nil {
		return err
	}
	fmt.Printf(".gitignore written to %q\n", filepath.Join(root, ".gitignore"))

	confYml, err := yaml.Marshal(&conf)
	if err != nil {
		return err
	}

	noteBook := filepath.Join(root, conf.DefaltNoteBook)
	if err = os.Mkdir(noteBook, os.ModePerm); err != nil && !os.IsExist(err) {
		return err
	}
	fmt.Printf("default notebook created at %q\n", noteBook)

	db := core.NotesDb{}
	if err = db.WriteJson(root, conf.DefaltNoteBook); err != nil {
		return err
	}
	fmt.Printf("db for %q created\n", noteBook)

	fmt.Printf("\n%s\nwas written to %q\n", confYml, configFile)

	return os.WriteFile(configFile, confYml, 0666)
}

func getPassword() (string, error) {
	fmt.Println("There isn't any way to change the password for now. So give a strong one")
	fmt.Print("Enter the password: ")
	password, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}
	fmt.Println() // Print a newline after the password input
	return string(password), nil
}
