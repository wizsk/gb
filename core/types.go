package core

import (
	"io/fs"
	"os"
)

type fileReader func(name string) ([]byte, error)
type fileWriter func(name string, data []byte, mod os.FileMode) error
type fileStat func(name string) (fs.FileInfo, error)
type fileRemove func(name string) error
type editorOpener func(fileName string, editor string) error
