package healthPinger

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/ahr-i/triton-agent/setting"
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
		log.Println("*** (ERROR) There is no scheduler.")

		os.Exit(1)
	}
}
