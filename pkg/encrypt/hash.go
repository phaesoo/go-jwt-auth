package encrypt

import (
	"crypto/sha256"
	"encoding/hex"
)

func EncryptSHA256(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}