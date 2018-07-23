package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/render"

	"github.com/smarthut/agent/device"
)

func (api *API) getSocket(w http.ResponseWriter, r *http.Request) {
	var p device.Payload
	if err := render.DecodeJSON(r.Body, &p); err != nil {
		fmt.Println(err)
	}

	s, err := api.device.Read(p.ID)
	if err != nil {
		log.Println(err)
	}

	render.JSON(w, r, s)
}

func (api *API) postSocket(w http.ResponseWriter, r *http.Request) {
	var p device.Payload
	if err := render.DecodeJSON(r.Body, &p); err != nil {
		fmt.Println(err)
	}

	if err := api.device.Write(p.ID, p.Status); err != nil {
		fmt.Println(err)
	}

	s, err := api.device.Read(p.ID)
	if err != nil {
		fmt.Println(err)
	}

	render.JSON(w, r, s)
}
