package handler

import (
	"log"
	"net/http"

	"github.com/ahr-i/triton-agent/tritonCommunicator"
	"github.com/gorilla/mux"
)

func (h *Handler) readyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	model := vars["model"]
	version := vars["version"]

	tritonStatus, err := tritonCommunicator.Ready(model, version)
	if err != nil {
		log.Println("** (ERROR)", err)
		rend.JSON(w, http.StatusBadRequest, nil)
		return
	}

	if tritonStatus {
		rend.JSON(w, http.StatusOK, nil)
	} else {
		rend.JSON(w, http.StatusBadRequest, nil)
	}
}
