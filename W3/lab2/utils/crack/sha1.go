package crack

import (
	"crypto/sha1"
	"encoding/hex"
)

func Sha1_hash(text string) string {
	hash := sha1.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}