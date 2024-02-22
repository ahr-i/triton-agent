package callback

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/ahr-i/triton-agent/setting"
	"github.com/ahr-i/triton-agent/src/logCtrlr"
)

type requestData struct {
	Port      string  `json:"port"`
	BurstTime float64 `json:"burst_time"`
	Id        string  `json:"id"`
	Model     string  `json:"model"`
	Version   string  `json:"version"`
}

/* This is the callback function sent to the scheduler. */
/* Sends the burst time, which is the completion time of the inference. */
func Callback(burstTime float64, provider string, model string, version string) {
	// If you are not using a scheduler, ignore this function.
	if !setting.SchedulerActive {
		return
	}

	// Incorporate the burst time into the JSON data.
	request := requestData{
		Port:      setting.ServerPort,
		BurstTime: burstTime,
		Id:        provider,
		Model:     model,
		Version:   version,
	}
	jsonData, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}

	// Send a request to the scheduler.
	url := fmt.Sprintf("http://%s/%s", setting.SchedulerUrl, "callback")
	resp, _ := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if resp == nil || resp.StatusCode != http.StatusOK {
		logCtrlr.Error(errors.New("there is no scheduler"))
		return
	}
	logCtrlr.Log("Sent a callback request to the scheduler.")
}
