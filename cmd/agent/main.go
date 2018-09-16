package main

import (
	"fmt"
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"

	"github.com/smarthut/agent/api"
	"github.com/smarthut/agent/conf"
	"github.com/smarthut/agent/device"
)

var (
	version = "master"
	commit  = "none"
	date    = "unknown"
)

var pollingTime time.Duration

func main() {
	var config conf.Configuration
	if err := envconfig.Process("agent", &config); err != nil {
		log.Println(err)
	}

	pollingTime = config.Device.PollingTime

	device, err := device.New(config.Device.Driver, config.Device.Host, config.Device.Password)
	if err != nil {
		log.Println(err)
	}

	api, err := api.New(device)

	go startPolling(device)

	l := fmt.Sprintf("%s:%d", config.Host, config.Port)
	log.Printf("Starting SmartHut Agent %s on %s\n", version, l)
	api.Start(l)
}

func startPolling(d device.Device) {
	for {
		if err := d.Fetch(); err != nil {
			log.Println(err)
		}
		<-time.After(pollingTime)
	}
}
