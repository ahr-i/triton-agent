package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

var rend *render.Render = render.New()

type Handler struct {
	http.Handler
}

func CreateHandler() *Handler {
	mux := mux.NewRouter()
	handler := &Handler{
		Handler: mux,
	}

	mux.HandleFunc("/ping", handler.PingHandler).Methods("GET") // Ping Check
	mux.HandleFunc("/{model:[a-z-_]+}/{version:[0-9]+}/infer", handler.InferHandler).Methods("POST")

	return handler
}
