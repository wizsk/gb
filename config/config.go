package config

import (
	"path/filepath"
)

// default stuff
const (
	ConfigFileName  = "config.yml"
	RootDirEnvName  = "GB_ROOT_DIR"
	EncExt          = ".md.enc"
	DecExt          = ".md"
	KeyFileName     = ".key"
	defatutNoteBook = "default"

	editor = "nvim"
)

type Config struct {
	RootDir        string `yaml:",omitempty"` // root dir
	encExt         string
	decExt         string
	DefaltNoteBook string
	Key            string `yaml:",omitempty"`
	Editor         string `yaml:"editor"`
}

func DefaultConf() Config {
	return Config{
		encExt:         EncExt,
		decExt:         DecExt,
		DefaltNoteBook: defatutNoteBook,
		Editor:         editor,
	}
}

// adds c.encExt with the name
func (c *Config) FullEncFilePath(n string) string {
	return filepath.Join(c.RootDir, c.DefaltNoteBook, n+c.encExt)
}

// adds c.encExt with the name
func (c *Config) FullDecFilePath(n string) string {
	return filepath.Join(c.RootDir, c.DefaltNoteBook, n+c.decExt)
}
