package tool

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
)

func EncoderSha256(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	sum := h.Sum(nil)
	return hex.EncodeToString(sum)
}

func Md5(data string) string {
	m := md5.New()
	io.WriteString(m, data)
	sum := m.Sum(nil)
	return fmt.Sprintf("%x", sum)
}

func Base64Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func Base64Decode(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}