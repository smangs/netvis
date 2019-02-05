package devices

import (
	validator "gopkg.in/go-playground/validator.v9"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

// Device type is used to create IP device objects.
type Device struct {
	ID          string   `validate:"required,alphanum,eq=40"` // should be a sha1 with 40 alphanumeric characters
	MACAddress  string   `validate:"mac"`                     // should be a valid mac address
	IPAddresses []string `validate:"ip"`                      // should be a valid IP address
	Hostnames   []string `validate:"hostname"`                // should be a valid hostname
	DeviceClass string   `validate:"required,alpha"`          // used to store the device class (e.g. mobile, server, desktop, etc)
	DataClass   []string `validate:"required,alpha"`          // used to store data classification of device.
	IDScore     int      `validate:"required,numeric"`        // used to rank confidence level of device identity.
	IFFScore    int      `validate:"required,numeric"`        // used to rank device as friend or foe
	RiskScore   int      `validate:"required,numeric"`        // used to rank device's risk level to the University
}
