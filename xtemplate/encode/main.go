package encode

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
)

func HmacSha512(key, src string) string {
	m := hmac.New(sha512.New, []byte(key))
	m.Write([]byte(src))
	return hex.EncodeToString(m.Sum(nil))
}
