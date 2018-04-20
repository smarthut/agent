package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/render"
	"github.com/smarthut/agent/model/device"
)

// APIHandler handles root api route
func APIHandler(w http.ResponseWriter, r *http.Request) {
	dev := device.Get()
	render.JSON(w, r, dev)
}

// SocketGetHandler handles Get device method
func SocketGetHandler(w http.ResponseWriter, r *http.Request) {
	dev := device.Get()
	var p device.Payload
	if err := render.DecodeJSON(r.Body, &p); err != nil {
		fmt.Println(err)
	}

	s, err := dev.Get(p.ID)
	if err != nil {
		log.Println(err)
	}

	render.JSON(w, r, s)
}

// SocketPostHandler handles Set device method
func SocketPostHandler(w http.ResponseWriter, r *http.Request) {
	dev := device.Get()
	var p device.Payload
	if err := render.DecodeJSON(r.Body, &p); err != nil {
		fmt.Println(err)
	}

	if err := dev.Set(p.ID, p.Status); err != nil {
		fmt.Println(err)
	}

	s, err := dev.Get(p.ID)
	if err != nil {
		fmt.Println(err)
	}

	render.JSON(w, r, s)
}
