package tritonCommunicator

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ahr-i/triton-agent/setting"
	"github.com/ahr-i/triton-agent/src/logCtrlr"
)

func GetRepositoryIndex() ([]byte, error) {
	logCtrlr.Log("Get model list")
	url := fmt.Sprintf("http://%s/v2/repository/index", setting.TritonUrl)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	logCtrlr.Log("Success! model list:")
	log.Println(string(bodyBytes))

	return bodyBytes, nil
}
