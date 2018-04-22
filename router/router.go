package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/smarthut/agent/handler"
)

// New initializes routes
func New() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/", handler.APIHandler)
		r.Route("/", func(r chi.Router) {
			r.Get("/socket", handler.SocketGetHandler)
			r.Post("/socket", handler.SocketPostHandler)
		})
	})

	return r
}
