package p2p

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/ahr-i/triton-agent/setting"
)

type RequestData struct {
	Address string `json:"address"`
}

func postAlive(address string) {
	jsonData, err := json.Marshal(RequestData{Address: address})
	if err != nil {
		panic(err)
	}

	resp, _ := http.Post("http://"+setting.SchedulerUrl+"/alive", "application/json", bytes.NewBuffer(jsonData))
	if resp == nil || resp.StatusCode != http.StatusOK {
		log.Println("There Is No Scheduler.")

		os.Exit(1)
	}
}
