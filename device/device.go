package device

import (
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
