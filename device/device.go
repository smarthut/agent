package device

import (
	"errors"
	"os"

	"github.com/smarthut/agent/device/laurent112"
	"github.com/smarthut/agent/device/megad328"
)

// Device provides interface for accessing the device
type Device interface {
	UpdateSockets() error            // updates all sockets
	Get(id int) (interface{}, error) // get information from socket
	Set(id, value int) error         // sends information to socket
	Len() int                        // get sockets num
}

// Factory ...
type Factory func() (Device, error)

// NewMegaD328Device ...
func NewMegaD328Device() (Device, error) {
	host, ok := os.LookupEnv("DEVICE_HOST")
	if !ok {
		return nil, errors.New("megad328: DEVICE_HOST is required for the MegaD328 connection")
	}

	pass, ok := os.LookupEnv("DEVICE_PASS")
	if !ok {
		return nil, errors.New("megad328: DEVICE_PASS is required for the MegaD328 connection")
	}

	return megad328.New(host, pass), nil
}

// NewLaurent112Device ...
func NewLaurent112Device() (Device, error) {
	host, ok := os.LookupEnv("DEVICE_HOST")
	if !ok {
		return nil, errors.New("laurent112: DEVICE_HOST is required for the MegaD328 connection")
	}

	return laurent112.New(host), nil
}

// Drivers ...
var Drivers = map[string]Factory{
	"megad328":   NewMegaD328Device,
	"laurent112": NewLaurent112Device,
}
