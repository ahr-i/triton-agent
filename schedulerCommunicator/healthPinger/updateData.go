package healthPinger

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/ahr-i/triton-agent/setting"
)

func UpdateTaskInfo_burstTime(burstTime float64, provider string, model string, version string) {
	if _, exists := model_info[provider+"@"+model]; !exists {
		model_info[provider+"@"+model] = make(map[string]TaskInfo)
	}

	prevTaskInfo := model_info[provider+"@"+model][version]
	var nextTaskInfo TaskInfo

	if prevTaskInfo.AverageInferenceTime == 0 {
		nextTaskInfo = TaskInfo{
			AverageInferenceTime: float32(burstTime),
			LoadedAmount:         prevTaskInfo.LoadedAmount,
		}
	} else {
		nextTaskInfo = TaskInfo{
			AverageInferenceTime: (float32(burstTime) + prevTaskInfo.AverageInferenceTime) / 2,
			LoadedAmount:         prevTaskInfo.LoadedAmount,
		}
	}

	model_info[provider+"@"+model][version] = nextTaskInfo
}

func UpdateTaskInfo_start(provider string, model string, version string) {
	prevTaskInfo := model_info[provider+"@"+model][version]

	model_info[provider+"@"+model][version] = TaskInfo{
		AverageInferenceTime: prevTaskInfo.AverageInferenceTime,
		LoadedAmount:         prevTaskInfo.LoadedAmount + 1,
	}
}
func UpdateTaskInfo_end(provider string, model string, version string) {
	prevTaskInfo := model_info[provider+"@"+model][version]

	model_info[provider+"@"+model][version] = TaskInfo{
		AverageInferenceTime: prevTaskInfo.AverageInferenceTime,
		LoadedAmount:         prevTaskInfo.LoadedAmount - 1,
	}
}

type requestData struct {
	Port    string `json:"port"`
	Id      string `json:"id"`
	Model   string `json:"model"`
	Version string `json:"version"`
}

func UpdateModel(provider string, model string, version string) {
	modelkey := provider + "@" + model
	if _, exists := model_info[modelkey]; !exists {
		model_info[modelkey] = make(map[string]TaskInfo)
	}

	model_info[modelkey][version] = TaskInfo{AverageInferenceTime: 0, LoadedAmount: 0}

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
	//Manager에게 전달
	log.Println("Manager에게 modelupdate 전달 :", request)
	_, err = http.Post("http://"+setting.ManagerUrl+"/model/update", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("modelupdate정보 Manager에게 전달 실패 :", err)
	}

}
