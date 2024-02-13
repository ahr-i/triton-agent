package tritonCommunicator

import (
	"io/ioutil"
	"net/http"

	"github.com/ahr-i/triton-agent/setting"
)

func GetRepositoryIndex() ([]byte, error) {
	url := "http://" + setting.TritonUrl + "/v2/repository/index"
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

	return bodyBytes, nil
}
