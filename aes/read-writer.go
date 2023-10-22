package aes

import (
	"bytes"
	"errors"
	"io"
)

func Encrypt(input io.Reader, output io.Writer, keyStr string) error {
	if keyStr == "" {
		return errors.New("key string is empty")
	}

	in, err := io.ReadAll(input)
	if err != nil {
		return err
	} else if len(in) == 0 {
		return errors.New("invalid input with len 0")
	}
	key := stringToKey(keyStr)
	encData, err := encrypt(in, key)
	if err != nil {
		return err
	}
	_, err = io.Copy(output, bytes.NewReader(encData))
	return err
}

func Decrypt(input io.Reader, output io.Writer, keyStr string) error {
	if keyStr == "" {
		return errors.New("key string is empty")
	}

	in, err := io.ReadAll(input)
	if err != nil {
		return err
	} else if len(in) == 0 {
		return errors.New("invalid input with len 0")
	}

	key := stringToKey(keyStr)
	encData, err := decrypt(in, key)
	if err != nil {
		return err
	}
	_, err = io.Copy(output, bytes.NewReader(encData))
	return err
}
