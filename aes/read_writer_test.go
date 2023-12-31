package aes

import (
	"bytes"
	"errors"
	"reflect"
	"testing"
)

func TestEDPass(t *testing.T) {
	tests := []struct {
		key string
		msg []byte
	}{
		{
			key: StringToHashHex("key"),
			msg: []byte("very important msg that no one else should see"),
		},
		{
			key: StringToHashHex("u wont ever know the keyyyy"),
			msg: []byte("very important msg that no one else should see, ya u can't"),
		},
	}

	for _, tt := range tests {
		enc := new(bytes.Buffer)
		err := Encrypt(bytes.NewReader(tt.msg), enc, HexToHash(tt.key))
		if err != nil {
			t.Error(err)
			t.FailNow()
		}

		dec := new(bytes.Buffer)
		err = Decrypt(enc, dec, HexToHash(tt.key))
		if err != nil {
			t.Error(err)
			t.FailNow()
		}

		if !reflect.DeepEqual(tt.msg, dec.Bytes()) {
			t.Error("expected", dec.String(), "to be", string(tt.msg))
			t.FailNow()
		}
	}
}

func TestEDFail(t *testing.T) {
	tests := []struct {
		key  []byte
		dkey []byte
		msg  []byte
		err  error // encyption err
		drr  error // decyption err
	}{
		// encryption err
		{
			key: []byte{},
			msg: []byte("very important msg that no one else should see"),
			err: errors.New("fo"),
		},
		{
			key: HexToHash(StringToHashHex("u wont ever know the keyyyy")),
			msg: []byte{},
			err: errors.New("fo"),
		},
		// decription err
		{
			key:  HexToHash(StringToHashHex("u wont ever know the keyyyy")),
			dkey: HexToHash(StringToHashHex("wrong key")),
			msg:  []byte("very important msg that no one else should see"),
			drr:  errors.New("fo"),
		},
	}

	for _, tt := range tests {
		enc := new(bytes.Buffer)
		err := Encrypt(bytes.NewReader(tt.msg), enc, tt.key)
		if tt.err == nil && err != nil {
			t.Errorf("error '%v' was expected to be nil", err)
			t.FailNow()
		}

		// skip decryption while the encrypion err is present
		if tt.err != nil {
			continue
		}

		dec := new(bytes.Buffer)
		err = Decrypt(enc, dec, tt.dkey)
		if tt.drr != nil && err == nil {
			t.Errorf("error '%v' was expected to be non-nil", err)
			t.FailNow()
		}

	}
}
