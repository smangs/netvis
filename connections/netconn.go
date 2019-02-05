package Connections

import (
	"crypto/sha256"
	"encoding/hex"
	"log"

	validator "gopkg.in/go-playground/validator.v9"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

// Connection type is used to store Connection objects.
// ID should be generated with Device.CreateSha256("IP" + "Port") function.
type Connection struct {
	ID   string `validate:"required"`
	IP   string `validate:"required,ip"`
	Port string `validate:"required"`
}

// createSha256 returns a Sha256 hash value for parameter "s".
func createSha256(s string) string {
	// get sha256 hash, convert and return [32]byte as string
	sum := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sum[:])
}

// CreateConnection is used to register a new observed connection for Device.
// It returns a new Connection object.
func CreateConnection(connIP string, connPort string) (Connection, error) {
	// TODO validate Port format "tcp/22, udp/53"

	// create Sha256 for Connection.ID
	connID := createSha256(connIP + connPort)

	// create initial Connection object.
	conn := Connection{
		ID:   connID,
		IP:   connIP,
		Port: connPort,
	}

	// validate the Connection struct
	validate = validator.New()
	err := validate.Struct(conn)
	if err != nil {
		log.Println(err)
		return Connection{}, err
	}

	return conn, nil
}
