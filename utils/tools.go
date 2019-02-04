package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func EncryptPassword(data []byte) string {
	hash := sha256.New()
	hash.Write(data)
	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)
	return mdStr
}
