package config

import (
	"path/filepath"
)

const (
	rootDirEnvName = "GB_ROOT_DIR"
	configFileName = "config.yml"
	defautRoot     = "GB_ROOT_DIR"
	encExt         = ".md.enc"
	decExt         = ".md"
	editor         = "nvim"
)

type Config struct {
	RootDir string `yaml:",omitempty"` // root dir
	encExt  string
	decExt  string
	Key     string `yaml:"key"`
	Editor  string `yaml:"editor"`
}

func DefaultConf() Config {
	return Config{
		encExt: encExt,
		decExt: decExt,
		Editor: editor,
		Key:    "",
	}
}

func (c *Config) AddEncExt(n string) string {
	return n + c.encExt
}

// adds c.encExt with the name
func (c *Config) FullEncFilePath(n string) string {
	return filepath.Join(c.RootDir, c.AddEncExt(n))
}
