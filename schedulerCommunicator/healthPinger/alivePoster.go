package healthPinger

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ahr-i/triton-agent/setting"
	"github.com/ahr-i/triton-agent/src/logCtrlr"
)

type RequestData struct {
	Port       string                         `json:"port"`
	Gpuname    string                         `json:"gpuname"`
	Model_info map[string]map[string]TaskInfo `json:"model_info"`
}

func postAlive() {
	jsonData, err := json.Marshal(RequestData{
		Port:       port,
		Gpuname:    gpuName,
		Model_info: model_info,
	})
	if err != nil {
		panic(err)
	}

	resp, _ := http.Post("http://"+setting.ManagerUrl+"/alive", "application/json", bytes.NewBuffer(jsonData))
	if resp == nil || resp.StatusCode != http.StatusOK {
		logCtrlr.Error(errors.New("there is no scheduler"))

		//os.Exit(1)
	}
}
