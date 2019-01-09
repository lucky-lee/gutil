package gEncrypt

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
)

//encrypty util

//hmac sha256 encrypt
func HmacSha256(content string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(content))

	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

//hmac sha1 encrypt
func HmacSha1(content string, secret string) string {
	h := hmac.New(sha1.New, []byte(secret))
	h.Write([]byte(content))

	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
