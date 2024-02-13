package tritonCommunicator

import (
	"fmt"
	"net/http"

	"github.com/ahr-i/triton-agent/setting"
	"github.com/ahr-i/triton-agent/src/logCtrlr"
)

func Ready(model string, version string) (bool, error) {
	logCtrlr.Log("Model check.")
	url := fmt.Sprintf("http://%s/v2/models/%s/versions/%s/ready", setting.TritonUrl, model, version)

	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}

	if resp.StatusCode == http.StatusOK {
		return true, nil
	} else {
		return false, nil
	}
}
