package callback

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ahr-i/triton-agent/schedulerCommunicator/healthPinger"
	"github.com/ahr-i/triton-agent/setting"
)

type requestData struct {
	Port    string `json:"port"`
	Id      string `json:"id"`
	Model   string `json:"model"`
	Version string `json:"version"`
}

/* This is the callback function sent to the scheduler. */
/* Sends the burst time, which is the completion time of the inference. */
func Callback(burstTime float64, provider string, model string, version string) {

	healthPinger.UpdateTaskInfo_burstTime(burstTime, provider, model, version)

	//Manager에게 이 모델에 대한 Inference 끝났다고 알림:w

	request := requestData{
		Port:    setting.ServerPort,
		Id:      provider,
		Model:   model,
		Version: version,
	}
	jsonData, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}
	url := fmt.Sprintf("http://%s/%s", setting.ManagerUrl, "inference/end")
	_, err = http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println(err)
	}

	// // If you are not using a scheduler, ignore this function.
	// if !setting.SchedulerActive {
	// 	return
	// }

	// // Incorporate the burst time into the JSON data.
	// request := requestData{
	// 	Port:      setting.ServerPort,
	// 	BurstTime: burstTime,
	// 	Id:        provider,
	// 	Model:     model,
	// 	Version:   version,
	// }
	// jsonData, err := json.Marshal(request)
	// if err != nil {
	// 	panic(err)
	// }

	// // Send a request to the scheduler.
	// url := fmt.Sprintf("http://%s/%s", setting.SchedulerUrl, "callback")
	// resp, _ := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	// if resp == nil || resp.StatusCode != http.StatusOK {
	// 	logCtrlr.Error(errors.New("there is no scheduler"))
	// 	return
	// }
	// logCtrlr.Log("Sent a callback request to the scheduler.")
}
