package tritonCommunicator

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ahr-i/triton-agent/setting"
)

/* Send a request to the Triton server and return the response. */
func Inference(model string, version string, request []byte) ([]byte, error) {
	log.Println("* (System) Request to triton server.")
	// URL setting.
	url := "http://" + setting.TritonUrl + "/v2/models/" + model + "/versions/" + version + "/infer"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(request))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	// Triton Server Response
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read body
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Println("* (System) Received an inference response from the Triton server.")

	return bodyBytes, nil
}
