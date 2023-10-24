package aes

import "errors"

func Enc(data, key []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, errors.New("Enc: key len is not 32 bytes")
	}

	if len(data) == 0 {
		return nil, errors.New("Enc: data len is 0 bytes")
	}

	return encrypt(data, key)
}

func Dec(data, key []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, errors.New("Dec: key len is not 32 bytes")
	}

	if len(data) == 0 {
		return nil, errors.New("Dec: data len is 0 bytes")
	}

	return decrypt(data, key)
}
