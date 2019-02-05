package Device

import (
	"crypto/sha256"
	"encoding/hex"

	validator "gopkg.in/go-playground/validator.v9"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

// Device type is used to create IP device objects.
type Device struct {
	ID          int      `validate:"required"`
	IPAddresses []string `validate:"required,ip"`
	Hostnames   []string `validate:"hostname"`
}

// createSha256 returns a Sha256 hash value for parameter "s".
func createSha256(s string) string {
	// get sha256 hash, convert and return [32]byte as string
	sum := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sum[:])
}
