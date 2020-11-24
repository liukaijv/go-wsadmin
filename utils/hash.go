package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5Str(plainText string) string {
	m := md5.New()
	m.Write([]byte(plainText))
	return hex.EncodeToString(m.Sum(nil))
}


