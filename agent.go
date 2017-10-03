package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/smarthut/agent/device"
)

var dev device.Device

func main() {
	os.Setenv("DEVICE_HOST", "192.168.1.102")
	os.Setenv("DEVICE_PASS", "sec")
	factory, ok := device.Drivers["megad328"]
	if !ok {
		fmt.Println("not ok")
	}

	var err error
	dev, err = factory()
	if err != nil {
		fmt.Println(err)
	}

	go loop()

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)

	r.Get("/", apiHandler)
	r.Get("/{socketID:^[0-9]*$}", socketGetHandler)
	r.Post("/{socketID:^[0-9]*$}", socketPostHandler)

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
	id, err := strconv.Atoi(chi.URLParam(r, "socketID"))
	if err != nil {
		log.Println(err)
		return
	}

	if id < 0 || id >= dev.Len() {
		w.Write([]byte("agent: id put of bounds"))
		return
	}

	s, err := dev.Get(id)
	if err != nil {
		log.Println(nil)
	}

	body, err := json.Marshal(s)
	if err != nil {
		log.Println("agent: unable to marshal device")
		return
	}
	w.Write(body)
}

func socketPostHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "socketID"))
	if err != nil {
		log.Println(err)
		return
	}

	if id < 0 || id >= dev.Len() {
		w.Write([]byte("agent: id put of bounds"))
		return
	}

	stringValue := r.FormValue("value")
	value, err := strconv.Atoi(stringValue)
	if err != nil {
		log.Printf("agent: unable to set %s to %d", stringValue, id)
	}

	if err = dev.Set(id, value); err != nil {
		log.Println(err)
	}

	s, err := dev.Get(id)
	if err != nil {
		log.Println(nil)
	}

	body, err := json.Marshal(s)
	if err != nil {
		log.Println("agent: unable to marshal device")
		return
	}
	w.Write(body)
}
