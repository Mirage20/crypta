package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/binary"
)

const (
	AesKeyLength = 32
)

func Encrypt(plainBytes []byte, pubKey *rsa.PublicKey) ([]byte, error) {

	randKey := make([]byte, AesKeyLength)
	if _, err := rand.Read(randKey); err != nil {
		return nil, err
	}

	rsaCipherBytes, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, pubKey, randKey, nil)
	if err != nil {
		return nil, err
	}

	cipherBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(cipherBytes, uint16(len(rsaCipherBytes)))
	cipherBytes = append(cipherBytes, rsaCipherBytes...)

	blockCipher, err := aes.NewCipher(randKey)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, err
	}

	zeroNonce := make([]byte, aesgcm.NonceSize())
	cipherBytes = aesgcm.Seal(cipherBytes, zeroNonce, plainBytes, nil)
	return cipherBytes, nil
}
