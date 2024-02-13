package handler

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ahr-i/triton-agent/src/httpController"
	"github.com/ahr-i/triton-agent/src/logCtrlr"
	"github.com/ahr-i/triton-agent/tritonCommunicator"
	"github.com/gorilla/mux"
)

func (h *Handler) inferHandler(w http.ResponseWriter, r *http.Request) {
	// Extract model information from the URL.
	vars := mux.Vars(r)
	model := vars["model"]
	version := vars["version"]

	// Extract the request from the body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logCtrlr.Error(err)
		return
	}
	logCtrlr.Log("Request: ▽▽▽▽▽▽▽▽▽▽")
	log.Println(string(body))

	// Request to triton
	response, err := tritonCommunicator.Inference(model, version, body)
	if err != nil {
		logCtrlr.Error(err)
		return
	}

	httpController.JSON(w, http.StatusOK, response)
}
