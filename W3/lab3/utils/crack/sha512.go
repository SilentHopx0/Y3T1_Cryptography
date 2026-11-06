package crack

import (
	"crypto/sha512"
	"encoding/hex"
)

func Sha1_hash(text string) string {
	hash := sha512.Sum512([]byte(text))
	return hex.EncodeToString(hash[:])
}