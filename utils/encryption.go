package utils

import (
	"crypto/sha256"
	"fmt"
)

func NewSHA256(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}
func EncryptString(s string) string {
	return fmt.Sprintf("%x",NewSHA256([]byte(s)))
}
