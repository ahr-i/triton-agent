package modelStoreCommunicator

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ahr-i/triton-agent/setting"
)

type modelStoreResponse struct {
	File string `json:"file"`
}

func GetModel(provider string, model string, version string) ([]byte, error) {
	// Code modifications are needed.
	// Url setting.
	url := fmt.Sprintf("http://%s/getModel?id=%s&filename=%s", setting.ModelStoreUrl, provider, model+".zip")

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response modelStoreResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	modelFile, err := base64.StdEncoding.DecodeString(response.File)
	if err != nil {
		return nil, err
	}

	return modelFile, nil
}
