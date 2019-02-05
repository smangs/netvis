package devices

import (
	"log"

	validator "gopkg.in/go-playground/validator.v9"
)

// use a single instance of Validate, it caches struct info
var valid *validator.Validate

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

// validateDevice func is used to validate the object
func validateDevice(d *Device) error {
	// TODO:
	return nil
}

// calcIDScore func is used to calculate the confidence level of the device indentity.
func (d *Device) calcIDScore() error {
	// TODO:
	d.IDScore = 1
	return nil
}

// calcIFFScore func is used to calculate the device threat score based on observed behavior or other information.
func (d *Device) calcIFFScore() error {
	// TODO:
	d.IFFScore = 1
	return nil
}

// calcRiskScore func is used to calculate the risk level of the device.
func (d *Device) calcRiskScore() error {
	// TODO:
	d.RiskScore = 1
	return nil
}

// Init func is used to initialize a Device object with calculated default values and validates the Device object.
func (d *Device) Init() error {
	err := d.calcIDScore()
	if err != nil {
		// not a critical error
		log.Println("ERROR - could not execute Device.calcIDScore()")
	}
	err2 := d.calcIFFScore()
	if err2 != nil {
		// not a critical error
		log.Println("ERROR - could not execute Device.calcIFFScore()")
	}
	err3 := d.calcRiskScore()
	if err3 != nil {
		// not a critical error
		log.Println("ERROR - could not execute Device.calcRiskScore()")
	}
	err4 := validateDevice(d)
	if err4 != nil {
		// not a critical error
		log.Println("ERROR - could not validate device while executing Device.Init()")
	}
	// device validation is a critical error
	return err4
}
