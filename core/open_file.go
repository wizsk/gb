package core

import (
	"os"
	"os/exec"

	"github.com/wizsk/gb/aes"
	"github.com/wizsk/gb/config"
)

const (
	readWritePermission = 0666
)

// filename should be the base file name. Like: "fo" nor "fo.md.enc"
func OpenFile(c *config.Config, fileName string) error {
	return open(os.ReadFile, os.WriteFile, os.Remove, openEditor, c, fileName)
}

func open(reader func(string) ([]byte, error),
	writer func(string, []byte, os.FileMode) error,
	remove func(string) error,
	oedit func(string, string) error, // open editor
	conf *config.Config, fileName string) error {

	encFile := conf.FullEncFilePath(fileName)
	decFile := conf.FullDecFilePath(fileName)

	// read the enc
	encData, err := reader(encFile)
	if err != nil {
		return err
	}

	// decdata
	decData, err := aes.Dec(encData, aes.HexToHash(conf.Key))
	if err != nil {
		return err
	}

	if err = writer(decFile, decData, readWritePermission); err != nil {
		return err
	}
	defer remove(decFile) // clean the file

	if err = oedit(decFile, conf.Editor); err != nil {
		return err
	}

	ChangedDecData, err := reader(decFile)
	if err != nil {
		return err
	}

	changedEncData, err := aes.Enc(ChangedDecData, aes.HexToHash(conf.Key))
	if err != nil {
		return err
	}

	if err = writer(encFile, changedEncData, readWritePermission); err != nil {
		return err
	}

	return nil
}

func openEditor(file, editor string) error {
	cmd := exec.Command(editor, file)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
