package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kelseyhightower/envconfig"

	"github.com/smarthut/agent/model/device"
	"github.com/smarthut/agent/router"
)

var (
	version = "master"
	commit  = "none"
	date    = "unknown"
)

func main() {
	var conf device.Configuration
	if err := envconfig.Process("agent", &conf); err != nil {
		log.Println(err)
	}

	if err := device.Create(conf); err != nil {
		log.Println(err)
	}

	l := fmt.Sprintf("%s:%d", conf.Host, conf.Port)

	log.Printf("Starting SmartHut Agent %s on %s\n", version, l)
	http.ListenAndServe(l, router.New())
}
