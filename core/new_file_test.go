package core

import (
	"bytes"
	"io/fs"
	"os"
	"reflect"
	"testing"

	"github.com/wizsk/gb/aes"
	"github.com/wizsk/gb/config"
)

func TestNewNote(t *testing.T) {
	enc := new(bytes.Buffer)
	dec := new(bytes.Buffer)
	conf := config.DefaultConf()
	conf.Key = aes.StringToHashHex("foo bar bazz")
	fN := "fo" // filename

	err := newNote(&conf, fN,
		func(_ string) ([]byte, error) {
			return dec.Bytes(), nil
		},
		func(_ string, b []byte, _ os.FileMode) error {
			_, err := enc.Write(b)
			return err
		},
		func(s string) (fs.FileInfo, error) {
			return nil, os.ErrNotExist
		},
		func(_ string) error {
			t.Log("remove was called")
			return nil
		},
		func(_, _ string) error {
			_, err := dec.WriteString("yo meo beoooooo")
			return err
		},
	)
	IsNil(t, err)

	decData, err := aes.Dec(enc.Bytes(), aes.HexToHash(conf.Key))
	IsNil(t, err)
	if !reflect.DeepEqual(decData, dec.Bytes()) {
		t.Error("file corruped")
		t.FailNow()
	}

}
