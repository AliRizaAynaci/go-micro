package main

import (
	"net/http"
)

func (app *Config) Broker(w http.ResponseWriter, req *http.Request) {
	payload := JsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}