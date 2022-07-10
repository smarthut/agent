package agent

import (
	"fmt"

	"github.com/smarthut/agent/driver/laurent112"
	"github.com/smarthut/agent/driver/megad328"
)

// Device implements abstract device accessor
type Device interface {
	Read(id int) (interface{}, error)
	Write(id int, status interface{}) error
	Ping() (bool, error)
	Fetch() error
}

// New creates the specified device
func New(driver, host, password string) (Device, error) {
	factory, ok := Drivers[driver]
	if !ok {
		return nil, fmt.Errorf("agent: there are no driver for %s device", driver)
	}

	dev, err := factory(host, password)
	if err != nil {
		return nil, err
	}

	return dev, nil
}

// Factory ...
type Factory func(host, password string) (Device, error)

func newMegaD328Device(host, password string) (Device, error) {
	return megad328.New(host, password), nil
}

func newLaurent112Device(host, password string) (Device, error) {
	return laurent112.New(host, password), nil
}

// Drivers ...
var Drivers = map[string]Factory{
	"megad328":   newMegaD328Device,
	"laurent112": newLaurent112Device,
}
