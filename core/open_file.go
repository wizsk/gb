package core

import (
	"fmt"
	"io/fs"
	"os"
	"time"

	"github.com/wizsk/gb/aes"
	"github.com/wizsk/gb/config"
)

// filename should be the base file name. Like: "fo" not "fo.md.enc"
func OpenFile(c *config.Config, fileName string) error {
	return open(c, fileName, os.ReadFile, os.WriteFile, os.Stat, os.Remove, openEditor)
}

func open(conf *config.Config, fileName string,
	reader func(string) ([]byte, error),
	writer func(string, []byte, os.FileMode) error,
	stat func(string) (fs.FileInfo, error),
	remove func(string) error,
	oedit func(string, string) error, // open editor
) error {

	encFile := conf.FullEncFilePath(fileName)
	decFile := conf.FullDecFilePath(fileName)

	if _, err := stat(decFile); err != nil && !os.IsNotExist(err) {
		return err
	} else if err == nil {
		return fmt.Errorf("open: %q already opened", fileName)
	}

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
	defer clean(remove, decFile) // clean the file

	done := make(chan struct{})
	go func(done <-chan struct{}) {
		tick := time.NewTicker(time.Second)
		var modTime time.Time

		for {
			select {
			case <-done:
				return
			case <-tick.C:
				// file mod time changed sooo file changed.
				if stat, err := os.Stat(decFile); err == nil && stat.ModTime() != modTime {
					data, _ := reader(decFile)
					data, _ = aes.Enc(data, aes.HexToHash(conf.Key))
					_ = writer(encFile, data, readWritePermission)
					modTime = stat.ModTime()
				}
			}
		}
	}(done)

	if err = oedit(decFile, conf.Editor); err != nil {
		return err
	}

	// after editing is done the go routine func shoudl stop
	done <- struct{}{}

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

func clean(remove func(string) error, file string) {
	if err := remove(file); err != nil {
		fmt.Printf("ERROR: while removing %q\nPLEASE REMOVE THE FILE MANUALLY OTHERWISE IT MAY CASE DATA LEAK!", file)
	}
}
