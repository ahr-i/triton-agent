package handler

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/ahr-i/triton-agent/schedulerCommunicator/callback"
	"github.com/ahr-i/triton-agent/src/httpController"
	"github.com/ahr-i/triton-agent/src/logCtrlr"
	"github.com/ahr-i/triton-agent/tritonCommunicator"
	"github.com/gorilla/mux"
)

var mutex sync.Mutex

func (h *Handler) inferV2Handler(w http.ResponseWriter, r *http.Request) {
	// Extract model information from the URL.
	vars := mux.Vars(r)
	provider := vars["provider"]
	model := vars["model"]
	version := vars["version"]

	//healthPinger.UpdateTaskInfo_start(provider, model, version)
	mutex.Lock()
	defer mutex.Unlock()

	// Extract the request from the body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logCtrlr.Error(err)
		rend.JSON(w, http.StatusBadRequest, nil)
		return
	}
	printModelInfo(provider, model, version, string(body))

	// // Set model repository
	// if err := tritonController.ChangeModelRepository(provider, model, version); err != nil {
	// 	logCtrlr.Error(err)
	// 	rend.JSON(w, http.StatusBadRequest, nil)
	// 	return
	// }

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
	//healthPinger.UpdateTaskInfo_end(provider, model, version)

	httpController.JSON(w, http.StatusOK, response)
	//httpController.JSON(w, http.StatusOK, nil)

}

func printModelInfo(provider string, model string, version string, request string) {
	logCtrlr.Log("Request: ▽▽▽▽▽▽▽▽▽▽")
	log.Println("Provider:", provider)
	log.Println("Model:", model)
	log.Println("Version:", version)
	log.Println("Request:", request)
}

func (h *Handler) testInferV2Handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	provider := vars["provider"]
	model := vars["model"]
	version := vars["version"]

	mutex.Lock()
	defer mutex.Unlock()

	//요청받으면 랜덤한 인퍼런스 타임으로 결과값 돌려주기.
	inferTime := getRandNum(500, 10000)
	time.Sleep(time.Millisecond * time.Duration(inferTime))

	//랜덤한 확률로 인퍼런스 중 fault상황
	randRate := rand.Float64()
	if randRate < 0.1 {
		os.Exit(1)
	}

	//정상 수행 후 응답 상황
	burstTime := float64(inferTime) / 1000
	callback.Callback(burstTime, provider, model, version)
}

func getRandNum(min int, max int) int {
	return rand.Intn(max-min+1) + min
}
