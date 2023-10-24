package core

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"reflect"
	"testing"

	"github.com/wizsk/gb/aes"
	"github.com/wizsk/gb/config"
)

func TestOpen(t *testing.T) {
	conf := config.DefaultConf()
	conf.RootDir = "."
	conf.Key = aes.StringToHashHex("foooo")
	fileName := "some-file-name"
	encFile := conf.FullEncFilePath(fileName)
	decFile := conf.FullDecFilePath(fileName)

	mc := new(bytes.Buffer)
	mc.WriteString("this is some note\n")

	fl := new(bytes.Buffer)
	e, err := aes.Enc(mc.Bytes(), aes.HexToHash(conf.Key))
	IsNil(t, err)
	fl.Write(e)

	// read encFile
	// decryt and wirte to decfile
	// read decFile
	// enctypt and wirte to encFile
	err = open(
		&conf,
		fileName,
		// read
		func(s string) ([]byte, error) {
			if s == encFile {
				return fl.Bytes(), nil
			} else if s == decFile {
				return mc.Bytes(), nil
			}
			return nil, fmt.Errorf("unexpected %q filename", s)
		},
		// write
		func(s string, b []byte, fm os.FileMode) error {
			if s == encFile {
				fl.Truncate(0)
				_, err := fl.Write(b)
				return err
			} else if s == decFile {
				return nil
			}
			return fmt.Errorf("unexpected %q filename", s)
		},
		// stat
		func(s string) (fs.FileInfo, error) {
			return nil, os.ErrNotExist
		},
		// remove
		func(s string) error {
			return nil
		},
		// edit
		func(_, editor string) error {
			mc.WriteString("fo is not good brah\n")
			return nil
		},
	)
	IsNil(t, err)

	d, _ := aes.Dec(fl.Bytes(), aes.HexToHash(conf.Key))
	if !reflect.DeepEqual(d, mc.Bytes()) {
		t.Error("file corrupted")
		t.FailNow()
	}
}

func IsNil(t *testing.T, v any) {
	t.Helper()
	if v != nil {
		t.Errorf("'%v' was expected to be nil", v)
		t.FailNow()
	}

}
