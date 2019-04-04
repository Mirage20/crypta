package base64

import "encoding/base64"

func Decode(data []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(data))
}
