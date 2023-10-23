package initConf

import (
	"os"

	"github.com/wizsk/gb/config"
)

func InitConf() error {
	root, err := config.RootDir()
	if err != nil {
		return err
	}

	// check if the file exists or not
	if stat, err := os.Stat(root); err == nil {
	} else if os.IsNotExist(err) {
		err := os.Mkdir(root)
		if err != nil {
			return err
		}

	}
}
