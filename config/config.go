package config

import (
	"os"
	"path/filepath"
)

const (
	rootDirEnvName = "GB_ROOT_DIR"
	encExt         = ".md.enc"
	decExt         = ".md"
	editor         = "nvim"
)

type Config struct {
	RootDir string // root dir
	encExt  string
	decExt  string
	Key     string
	Editor  string
}

func DefaultConf() (*Config, error) {
	var root string
	if root = os.Getenv(rootDirEnvName); root == "" {
		var err error
		root, err = os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		root = filepath.Join(root, ".gb")
	}

	return &Config{
		RootDir: filepath.Join(),
		encExt:  encExt,
		decExt:  decExt,
		Editor:  editor,
		Key:     "",
	}, nil
}

func (c *Config) AddEncExt(n string) string {
	return n + c.encExt
}

// adds c.encExt with the name
func (c *Config) FullEncFilePath(n string) string {
	return filepath.Join(c.RootDir, c.AddEncExt(n))
}
