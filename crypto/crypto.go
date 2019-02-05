package Crypto

import (
	"crypto/sha256"
	"encoding/hex"
)

// CreateSha256 returns a Sha256 hash value for parameter "s".
func CreateSha256(s string) string {
	// get sha256 hash, convert and return [32]byte as string
	sum := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sum[:])
}
