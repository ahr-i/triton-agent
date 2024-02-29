package handler

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/ahr-i/triton-agent/schedulerCommunicator/callback"
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
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logCtrlr.Error(err)
		rend.JSON(w, http.StatusBadRequest, nil)
		return
	}
	logCtrlr.Log("Request: ▽▽▽▽▽▽▽▽▽▽")
	log.Println(string(body))

	// Request to tritons
	startTime := time.Now()
	response, err := tritonCommunicator.Inference(model, version, body)
	if err != nil {
		logCtrlr.Error(err)
		rend.JSON(w, http.StatusBadRequest, nil)
		return
	}
	endTime := time.Now()

	// Send burst time to scheduler
	burstTime := float64(endTime.Sub(startTime).Milliseconds()) / 1000
	log.Printf("* (SYSTEM) Burst time: %f\n", burstTime)
	callback.Callback(burstTime, "", model, version)

	httpController.JSON(w, http.StatusOK, response)
}
