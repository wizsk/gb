package aes

import (
	"errors"
	"reflect"
	"testing"
)

func TestAESPass(t *testing.T) {
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

	// passing tests
	for _, ts := range tests {
		t.Logf("testing key: %x,\nmsg: %s", ts.key, ts.msg)
		encrypted, err := encrypt(ts.msg, HexToHash(ts.key))
		if err != nil {
			t.Error(err)
			t.FailNow()
		}

		decrypted, err := decrypt(encrypted, HexToHash(ts.key))
		if err != nil {
			t.Error(err)
			t.FailNow()
		}

		if !reflect.DeepEqual(ts.msg, decrypted) {
			t.Error("msg & decriped msg don't match")
			t.FailNow()
		}
	}
}

func TestAESFail(t *testing.T) {
	tests := []struct {
		key, dkey string
		msg       []byte
		err       error // encryption err
		drr       error // decryption err
	}{
		{
			key: StringToHashHex("hhhh"),
			msg: []byte("very important msg that no one else should see"),
			err: errors.New("fo"),
		},
		{
			// empty input
			key: StringToHashHex("keyyyy"),
			msg: []byte{},
			err: errors.New("fo"),
		},
		{
			// empty input
			key: StringToHashHex("fo"),
			msg: []byte{},
			err: errors.New("fo"),
		},
		{
			// wrong d key
			key:  StringToHashHex("fo"),
			dkey: StringToHashHex("foxxxx"),
			msg:  []byte("yo hi"),
			drr:  errors.New("fo"),
		},
	}

	// passing tests
	for _, ts := range tests {
		t.Logf("testing key: %x,\nmsg: %s", ts.key, ts.msg)
		encrypted, err := encrypt(ts.msg, HexToHash(ts.key))
		// encryption should not fail
		if ts.err == nil && err != nil {
			t.Errorf("error '%v' was expected to be nil", err)
			t.FailNow()
		}

		// skip decryption while the encrypion err is present
		if ts.err != nil {
			continue
		}

		_, err = decrypt(encrypted, HexToHash(ts.dkey))
		if ts.drr != nil && err == nil {
			t.Errorf("error '%v' was expected to be non-nil", err)
			t.FailNow()
		}
	}
}
