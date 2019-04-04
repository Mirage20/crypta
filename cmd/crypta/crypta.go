package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mirage20/crypta/pkg/base64"
	"github.com/mirage20/crypta/pkg/crypto"
	"github.com/mirage20/crypta/pkg/io"
	"github.com/mirage20/crypta/pkg/rsa"
)

var (
	inFile        string
	outFile       string
	keyFile       string
	decryptMode   bool
	base64Encoded bool
	version       bool
)

var versionString string

func main() {
	flag.Parse()

	if version {
		fmt.Println("crypta", versionString)
		return
	}

	inData, err := io.ReadInput(inFile)
	check(err)
	if decryptMode {
		if base64Encoded {
			inData, err = base64.Decode(inData)
			check(err)
		}
		key, err := rsa.LoadPrivateKey(keyFile, "CRYPTA_PRIVATE_KEY")
		check(err)
		outData, err := crypto.Decrypt(inData, key)
		check(err)
		err = io.WriteOutput(outData, outFile)
		check(err)
	} else {
		key, err := rsa.LoadPublicKey(keyFile, "CRYPTA_PUBLIC_KEY")
		check(err)
		outData, err := crypto.Encrypt(inData, key)
		check(err)
		if base64Encoded {
			outData = base64.Encode(outData)
		}
		err = io.WriteOutput(outData, outFile)
		check(err)
	}
}

func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	flag.BoolVar(&base64Encoded, "base64", false, "Use base64 encoding to write output (encrypt mode), and base64 decoding to read input (decrypt mode)")
	flag.BoolVar(&decryptMode, "d", false, "Decrypt mode")
	flag.StringVar(&inFile, "in", "", "Input file path. Uses standard input if not provided")
	flag.StringVar(&keyFile, "key", "", "Public key file path for encrypting or private key file path for decrypting")
	flag.StringVar(&outFile, "out", "", "Output file path. Uses standard output if not provided")
	flag.BoolVar(&version, "version", false, "output version information")
}
