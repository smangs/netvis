package crypto

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

// CreateSha256 returns a Sha256 hash value for parameter "s".
func CreateSha256(s string) string {
	// get sha256 hash, convert and return [32]byte as string
	sum := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sum[:])
}

// CreateSha1 returns a Sha1 hash value for parameter "s".
func CreateSha1(s string) string {
	// get sha1 hash, convert and return byte as string
	sum := sha1.Sum([]byte(s))
	return hex.EncodeToString(sum[:])

}
