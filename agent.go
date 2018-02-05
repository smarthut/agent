package main

import (
	"fmt"
	"net/http"

	"github.com/smarthut/agent/model/device"
	"github.com/smarthut/agent/router"
)

func main() {
	if err := device.Create(); err != nil {
		fmt.Println(err)
	}

	http.ListenAndServe(":8080", router.New())
}
