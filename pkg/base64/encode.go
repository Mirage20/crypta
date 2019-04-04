package base64

import "encoding/base64"

func Encode(data []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(data))
}
