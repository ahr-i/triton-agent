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

	mux.HandleFunc("/ping", handler.pingHandler).Methods("GET") // Ping Check
	mux.HandleFunc("/{model:[a-z-_]+}/{version:[0-9]+}/infer", handler.inferHandler).Methods("POST")
	mux.HandleFunc("/repository/index", handler.repositoryIndexHandler).Methods("POST")
	mux.HandleFunc("/register/token", handler.registerTokenHandler).Methods("POST")
	mux.HandleFunc("/request/{model:[a-z-_]+}/{version:[0-9]+}/infer", handler.validateTokenAndInferenceHandler).Methods("POST")

	return handler
}
