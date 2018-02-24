package main

import (
	"fmt"
	"net/http"

	"github.com/kelseyhightower/envconfig"

	"github.com/smarthut/agent/model/device"
	"github.com/smarthut/agent/router"
)

func main() {
	var conf device.Configuration
	if err := envconfig.Process("agent", &conf); err != nil {
		fmt.Println(err)
	}

	if err := device.Create(conf); err != nil {
		fmt.Println(err)
	}

	l := fmt.Sprintf("%s:%d", conf.Host, conf.Port)

	http.ListenAndServe(l, router.New())
}
