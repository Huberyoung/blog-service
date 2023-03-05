package util

import (
	"crypto/md5"
	"encoding/hex"
)

func EncodeMd5(v string) string {
	m := md5.New()
	m.Write([]byte(v))
	return hex.EncodeToString(m.Sum(nil))
}
