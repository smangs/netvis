package Connections

import (
	"log"

	Crypto "github.com/smangs/netvis/crypto"
	validator "gopkg.in/go-playground/validator.v9"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

// NetConn type is used to store Connection objects.
// ID should be generated with Device.CreateSha256("IP" + "Port") function.
type NetConn struct {
	ID   string `validate:"required"`
	IP   string `validate:"required,ip"`
	Port string `validate:"required"`
}

// CreateNetConn is used to register a new observed connection for Device.
// It returns a new NetConn object.
func CreateNetConn(connIP string, connPort string) (NetConn, error) {
	// TODO validate Port format "tcp/22, udp/53"

	// create Sha256 for Connection.ID using netvis.crypto package
	connID := Crypto.CreateSha256(connIP + connPort)

	// create initial Connection object.
	conn := NetConn{
		ID:   connID,
		IP:   connIP,
		Port: connPort,
	}

	// validate the Connection struct
	validate = validator.New()
	err := validate.Struct(conn)
	if err != nil {
		log.Println(err)
		return NetConn{}, err
	}

	return conn, nil
}
