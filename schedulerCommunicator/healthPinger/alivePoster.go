package healthPinger

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/ahr-i/triton-agent/setting"
	"github.com/ahr-i/triton-agent/src/logCtrlr"
)

type RequestData struct {
	Port string `json:"port"`
}

func postAlive() {
	jsonData, err := json.Marshal(RequestData{Port: port})
	if err != nil {
		panic(err)
	}

	resp, _ := http.Post("http://"+setting.SchedulerUrl+"/alive", "application/json", bytes.NewBuffer(jsonData))
	if resp == nil || resp.StatusCode != http.StatusOK {
		logCtrlr.DError(errors.New("There is no scheduler."))

		os.Exit(1)
	}
}
