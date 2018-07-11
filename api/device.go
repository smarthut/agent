package api

import (
	"net/http"

	"github.com/go-chi/render"
)

func (api *API) getDevice(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, api.device)
}
