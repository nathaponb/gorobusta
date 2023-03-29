package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sha256Hasher(s string) string {

	hasher := sha256.New()
	hasher.Write([]byte(s))
	hashed := hasher.Sum(nil)
	hexStr := hex.EncodeToString(hashed)

	return hexStr
}
