package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
)

func Decrypt(cipherBytes []byte, key *rsa.PrivateKey) ([]byte, error) {
	if len(cipherBytes) < 2 {
		return nil, fmt.Errorf("cipher data does not contain rsa length information")
	}
	rsaCipherLength := int(binary.BigEndian.Uint16(cipherBytes))
	if len(cipherBytes) < rsaCipherLength+2 {
		return nil, fmt.Errorf("invalid rsa cipher length")
	}

	rsaCipherBytes := cipherBytes[2 : rsaCipherLength+2]
	aesCipherBytes := cipherBytes[rsaCipherLength+2:]

	aesKey, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, key, rsaCipherBytes, nil)
	if err != nil {
		return nil, err
	}

	blockCipher, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, err
	}

	zeroNonce := make([]byte, aesgcm.NonceSize())

	plaintext, err := aesgcm.Open(nil, zeroNonce, aesCipherBytes, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}
