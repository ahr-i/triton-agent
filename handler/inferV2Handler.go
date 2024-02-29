package handler

import (
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/ahr-i/triton-agent/schedulerCommunicator/callback"
	"github.com/ahr-i/triton-agent/src/httpController"
	"github.com/ahr-i/triton-agent/src/logCtrlr"
	"github.com/ahr-i/triton-agent/tritonCommunicator"
	"github.com/ahr-i/triton-agent/tritonController"
	"github.com/gorilla/mux"
)

var mutex sync.Mutex

func (h *Handler) inferV2Handler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	// Extract model information from the URL.
	vars := mux.Vars(r)
	provider := vars["provider"]
	model := vars["model"]
	version := vars["version"]

	// Extract the request from the body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logCtrlr.Error(err)
		rend.JSON(w, http.StatusBadRequest, nil)
		return
	}
	printModelInfo(provider, model, version, string(body))

	// Set model repository
	if err := tritonController.ChangeModelRepository(provider, model, version); err != nil {
		logCtrlr.Error(err)
		rend.JSON(w, http.StatusBadRequest, nil)
		return
	}

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
	callback.Callback(burstTime, provider, model, version)

	httpController.JSON(w, http.StatusOK, response)
}

func printModelInfo(provider string, model string, version string, request string) {
	logCtrlr.Log("Request: ▽▽▽▽▽▽▽▽▽▽")
	log.Println("Provider:", provider)
	log.Println("Model:", model)
	log.Println("Version:", version)
	log.Println("Request:", request)
}
