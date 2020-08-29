package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func ComputeHmac256(message string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	s := hex.EncodeToString(h.Sum(nil))
	fmt.Println(s)
	fmt.Println(base64.StdEncoding.EncodeToString(h.Sum(nil)))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
	//return s
}