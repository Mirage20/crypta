package rsa

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
)

func LoadPublicKey(file string, envVar string) (*rsa.PublicKey, error) {
	b, err := loadPem(file, envVar)
	if err != nil {
		return nil, err
	}
	pub, err := x509.ParsePKIXPublicKey(b)
	if err != nil {
		return nil, err
	}
	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("key is not a rsa public key")
	}
	return rsaPub, nil
}

func LoadPrivateKey(file string, envVar string) (*rsa.PrivateKey, error) {
	b, err := loadPem(file, envVar)
	if err != nil {
		return nil, err
	}
	return x509.ParsePKCS1PrivateKey(b)
}

func loadPem(file string, envVar string) ([]byte, error) {
	pemFile := ""
	if len(file) > 0 {
		pemFile = file
	} else {
		var ok bool
		pemFile, ok = os.LookupEnv(envVar)
		if !ok {
			return nil, fmt.Errorf("key file is not specified. Use -key or %s envarment varible to set the key file", envVar)
		}
	}
	fileBytes, err := ioutil.ReadFile(pemFile)
	if err != nil {
		return nil, err
	}
	pemBlock, _ := pem.Decode(fileBytes)
	return pemBlock.Bytes, nil
}
