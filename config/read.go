package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func Readconf() (*Config, error) {
	root, err := RootDir()
	if err != nil {
		return nil, err
	}

	if stat, err := os.Stat(root); err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("readconf: home dir for gb %q do not exist", root)
		} else {
			return nil, err
		}
	} else if !stat.IsDir() {
		return nil, fmt.Errorf("readconf: home dir for gb %q is not a dir", root)
	}

	confFile, err := os.ReadFile(filepath.Join(root, ConfigFileName))
	if err != nil {
		return nil, err
	}

	conf := DefaultConf()
	if err = yaml.Unmarshal(confFile, &conf); err != nil {
		return nil, err
	}

	return &conf, nil
}

func RootDir() (string, error) {
	var root string
	if root = os.Getenv(RootDirEnvName); root == "" {
		var err error
		root, err = os.UserHomeDir()
		if err != nil {
			return "", err
		}

		root = filepath.Join(root, ".gb")
	}
	return root, nil
}
