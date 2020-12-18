package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// md5加密
func GetMd5String(s string) string  {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}