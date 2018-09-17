package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/smarthut/agent"
)

// API is the main REST API
type API struct {
	handler http.Handler
	device  agent.Device
}

// New instantiates a new REST API
func New(device agent.Device) (*API, error) {
	api := &API{
		device: device,
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/", api.getDevice)
		r.Route("/", func(r chi.Router) {
			r.Get("/socket", api.getSocket)
			r.Post("/socket", api.postSocket)
		})
	})

	api.handler = r

	return api, nil
}

// Start starts API at address
func (api *API) Start(addr string) {
	http.ListenAndServe(addr, api.handler)
}
