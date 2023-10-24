package config

import (
	"path/filepath"
)

const (
	ConfigFileName = "config.yml"
	RootDirEnvName = "GB_ROOT_DIR"
	EncExt         = ".md.enc"
	DecExt         = ".md"

	editor = "nvim"
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
		encExt: EncExt,
		decExt: DecExt,
		Editor: editor,
	}
}

func (c *Config) AddEncExt(n string) string {
	return n + c.encExt
}

// adds c.encExt with the name
func (c *Config) FullEncFilePath(n string) string {
	return filepath.Join(c.RootDir, n+c.encExt)
}

// adds c.encExt with the name
func (c *Config) FullDecFilePath(n string) string {
	return filepath.Join(c.RootDir, n+c.decExt)
}
