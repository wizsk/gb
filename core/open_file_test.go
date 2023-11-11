package core

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/wizsk/gb/aes"
	"github.com/wizsk/gb/config"
)

type file struct {
	data *bytes.Buffer
	last time.Time
}

var _ fileInfo = &file{} // checking if it satisfies the interface

func (f *file) ModTime() time.Time {
	return f.last
}

func (f *file) Size() int64 {
	return int64(f.data.Len())
}

func TestOpen(t *testing.T) {
	conf := config.DefaultConf()
	conf.RootDir = "."
	conf.Key = aes.StringToHashHex("foooo")
	fileName := "some-file-name"
	encFile := conf.FullEncFilePath(fileName)
	decFile := conf.FullDecFilePath(fileName)

	// main content
	// unencrypted
	// dec := &file{new(bytes.Buffer), time.Now()}
	var dec *file = nil
	// dec.WriteString("this is some note\n")

	//  file
	// encrypted file should have something in
	// it for usually.
	var fl *file = &file{new(bytes.Buffer), time.Now()}
	d, err := aes.Enc([]byte("some abbettry datas"), aes.HexToHash(conf.Key))
	IsNil(t, err)
	fl.data.Write(d)

	// read encFile
	// decryt and wirte to decfile
	// read decFile
	// enctypt and wirte to encFile
	err = open(
		&conf, fileName,
		// read
		func(s string) ([]byte, error) {
			if s == encFile {
				return fl.data.Bytes(), nil
			} else if s == decFile {
				return dec.data.Bytes(), nil
			}
			return nil, fmt.Errorf("unexpected %q filename", s)
		},

		// write
		func(s string, b []byte, fm os.FileMode) error {
			if s == encFile {
				fl.data.Truncate(0)
				fl.last = time.Now()
				_, err := fl.data.Write(b)
				return err
			} else if s == decFile {
				// this should be called be
				return nil
			}
			return fmt.Errorf("unexpected %q filename", s)
		},
		// stat
		func(s string) (fileInfo, error) {
			if s == decFile {
				if dec == nil {
					return nil, os.ErrNotExist
				}
				return dec, nil
			}
			return nil, nil
		},
		// remove
		func(s string) error {
			return nil
		},
		// edit
		func(_, editor string) error {
			// creaing the file
			if dec == nil {
				dec = &file{new(bytes.Buffer), time.Now()}
			}
			dec.last = time.Now()
			_, err := dec.data.WriteString("fo is not good brah\n")
			return err

		},
	)

	IsNil(t, err)

	d, _ = aes.Dec(fl.data.Bytes(), aes.HexToHash(conf.Key))
	if !reflect.DeepEqual(d, dec.data.Bytes()) {
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
