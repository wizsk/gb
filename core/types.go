package core

import (
	"os"
	"time"
)

type fileInfo interface {
	ModTime() time.Time // modification time
	Size() int64        // length in bytes for regular files; system-dependent for others
}

type fileReader func(name string) ([]byte, error)
type fileWriter func(name string, data []byte, mod os.FileMode) error
type fileStat func(name string) (fileInfo, error)
type fileRemove func(name string) error
type editorOpener func(fileName string, editor string) error
