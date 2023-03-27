package main

import (
	"crypto/sha256"
	"encoding/hex"
)

func (app *Config) hasher(s string) string {

	h := sha256.New()
	h.Write([]byte(s))
	hashed := h.Sum(nil)
	hexStr := hex.EncodeToString(hashed)

	return hexStr
}
