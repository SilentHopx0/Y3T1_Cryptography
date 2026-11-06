package crack

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5_hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}