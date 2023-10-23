package aes

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
)

// Encrypt takes input, output as io.Reader, io.Writer
// and reads form the input and encrypts it then writes to output
//
// key []byte is a 32 bytes long key
func Encrypt(input io.Reader, output io.Writer, key []byte) error {
	if len(key) != 32 {
		return errors.New("Encrypt: key len is not 32 bytes")
	}

	in, err := io.ReadAll(input)
	if err != nil {
		return err
	} else if len(in) == 0 {
		return errors.New("invalid input with len 0")
	}

	encData, err := encrypt(in, key)
	if err != nil {
		return err
	}
	_, err = io.Copy(output, bytes.NewReader(encData))
	return err
}

// Decrypt takes input, output as io.Reader, io.Writer
// and reads form the input and decrypts it then writes to output
//
// key []byte is a 32 bytes long key
func Decrypt(input io.Reader, output io.Writer, key []byte) error {
	if len(key) != 32 {
		return errors.New("Decrypt: key len is not 32 bytes")
	}

	in, err := io.ReadAll(input)
	if err != nil {
		return err
	} else if len(in) == 0 {
		return errors.New("invalid input with len 0")
	}

	encData, err := decrypt(in, key)
	if err != nil {
		return err
	}
	_, err = io.Copy(output, bytes.NewReader(encData))
	return err
}

// StringToHash takes in a stirng
// return sha256sum of the string
//
// if the len of the string is 0 then it pannics
// func StringToHash(s string) []byte {
// 	if len(s) == 0 {
// 		panic("StringToHash: input can not be empty")
// 	}
// 	b32 := sha256.Sum256([]byte(s))
// 	return b32[:]
// }

// StringToHashHex takes in a stirng
// return sha256sum of the string and retuns the hex
//
// if the len of the string is 0 then it pannics
func StringToHashHex(s string) string {
	if len(s) == 0 {
		panic("StringToHash: input can not be empty")
	}

	b32 := sha256.Sum256([]byte(s))
	return hex.EncodeToString(b32[:])
}

// HexToHash takes in a stirng in hex
// return []byte of the string
//
// if errs then it pannics
func HexToHash(s string) []byte {
	if len(s) == 0 {
		panic("HexToHash: input can not be empty")
	}

	b, err := hex.DecodeString(s)
	if err != nil {
		panic("HexToHash: " + err.Error())
	}

	return b
}
