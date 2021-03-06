package main

import (
	"fmt"
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"

	"github.com/smarthut/agent"
	"github.com/smarthut/agent/api"
)

var (
	version = "master"
	commit  = "none"
	date    = "unknown"
)

func main() {
	var config agent.Configuration
	if err := envconfig.Process("agent", &config); err != nil {
		log.Println(err)
	}

	device, err := agent.New(config.Device.Driver, config.Device.Host, config.Device.Password)
	if err != nil {
		log.Println(err)
	}

	api, err := api.New(device)

	go startPolling(device, config.Device.PollingTime)

	l := fmt.Sprintf("%s:%d", config.Host, config.Port)
	log.Printf("Starting SmartHut Agent %s on %s\n", version, l)
	api.Start(l)
}

func startPolling(d agent.Device, t time.Duration) {
	for {
		if err := d.Fetch(); err != nil {
			log.Println(err)
		}
		<-time.After(t)
	}
}
