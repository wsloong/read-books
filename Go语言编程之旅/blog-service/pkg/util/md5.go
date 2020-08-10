package util

import (
	"crypto/md5"
	"encoding/hex"
)

func EncodeMD5(values string) string {
	m := md5.New()
	m.Write([]byte(values))
	return hex.EncodeToString(m.Sum(nil))
}
