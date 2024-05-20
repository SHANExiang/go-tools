package encrypt

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}