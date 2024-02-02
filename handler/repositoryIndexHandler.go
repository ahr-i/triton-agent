package handler

import (
	"io/ioutil"
	"net/http"

	"github.com/ahr-i/triton-agent/setting"
)

func (h *Handler) repositoryIndexHandler(w http.ResponseWriter, r *http.Request) {
	modelInformation, err := GetRepositoryIndex(setting.TritonUrl)
	if err != nil {
		panic(err)
	}

	rend.JSON(w, http.StatusOK, modelInformation)
}

func GetRepositoryIndex(tritonURL string) (string, error) {
	url := "http://" + tritonURL + "/v2/repository/index"
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}
