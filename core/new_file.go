package core

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/wizsk/gb/aes"
	"github.com/wizsk/gb/config"
)

func NewNote(conf *config.Config, fileName string) error {
	return newNote(conf, fileName, os.ReadFile, os.WriteFile, os.Stat, os.Remove, openEditor)
}

func newNote(conf *config.Config, fileName string,
	reader func(string) ([]byte, error),
	writer func(string, []byte, os.FileMode) error,
	stat func(string) (fs.FileInfo, error),
	remove func(string) error,
	oedit func(string, string) error, // open editor

) error {
	decFile := conf.FullDecFilePath(fileName)
	encFile := conf.FullEncFilePath(fileName)

	if _, err := stat(encFile); err != nil && !os.IsNotExist(err) {
		return err
	} else if err == nil {
		return fmt.Errorf("newNote: %q already exists", fileName)
	}

	if _, err := stat(decFile); err != nil && !os.IsNotExist(err) {
		return err
	} else if err == nil {
		return fmt.Errorf("newNote: %q already opened", fileName)
	}

	// editor can create the file
	if err := oedit(decFile, conf.Editor); err != nil {
		return err
	}
	defer clean(remove, decFile)

	decData, err := reader(decFile)
	if err != nil {
		return err
	}
	// clean the file

	if len(decData) == 0 {
		return fmt.Errorf("newNote: no data was written to %q", decFile)
	}

	encData, err := aes.Enc(decData, aes.HexToHash(conf.Key))
	if err != nil {
		return err
	}

	return writer(encFile, encData, readWritePermission)
}
