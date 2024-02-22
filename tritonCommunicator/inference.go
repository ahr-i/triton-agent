package tritonCommunicator

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ahr-i/triton-agent/setting"
	"github.com/ahr-i/triton-agent/src/logCtrlr"
)

/* Send a request to the Triton server and return the response. */
func Inference(model string, version string, request []byte) ([]byte, error) {
	logCtrlr.Log("Request to triton server.")
	// URL setting.
	url := fmt.Sprintf("http://%s/v2/models/%s/versions/%s/infer", setting.TritonUrl, model, version)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(request))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read body
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	logCtrlr.Log("Received an inference response from the Triton server.")

	return bodyBytes, nil
}
