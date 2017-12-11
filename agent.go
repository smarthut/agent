package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/caarlos0/env"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"

	"github.com/smarthut/agent/device"
)

var dev device.Device

func main() {
	var deviceConfig DeviceConfig
	err := env.Parse(&deviceConfig)
	if err != nil {
		fmt.Println(err)
	}

	dev, err = device.NewDevice(deviceConfig.Driver, deviceConfig.Host, deviceConfig.Password)
	if err != nil {
		fmt.Println(err)
	}

	go loop()

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)

	r.Get("/", apiHandler)

	r.Get("/socket", socketGetHandler)
	r.Post("/socket", socketPostHandler)

	http.ListenAndServe(":8080", r)
}

func loop() {
	for {
		// TODO: check if code needs some mutex
		if err := dev.UpdateSockets(); err != nil {
			log.Println(err)
		}
		<-time.After(5 * time.Second)
	}
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	body, err := json.Marshal(dev)
	if err != nil {
		log.Println("agent: unable to marshal device")
		return
	}
	w.Write(body)
}

func socketGetHandler(w http.ResponseWriter, r *http.Request) {
	var p device.Payload
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		fmt.Println(err)
	}

	if p.ID < 0 || p.ID >= dev.Len() {
		w.Write([]byte("agent: id put of bounds"))
		return
	}

	s, err := dev.Get(p.ID)
	if err != nil {
		log.Println(err)
	}

	render.JSON(w, r, s)
}

func socketPostHandler(w http.ResponseWriter, r *http.Request) {
	var p device.Payload
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		fmt.Println(err)
	}

	if p.ID < 0 || p.ID >= dev.Len() {
		w.Write([]byte("agent: id put of bounds"))
		return
	}

	if err = dev.Set(p.ID, p.Value); err != nil {
		log.Println(err)
	}

	s, err := dev.Get(p.ID)
	if err != nil {
		log.Println(err)
	}

	render.JSON(w, r, s)
}
