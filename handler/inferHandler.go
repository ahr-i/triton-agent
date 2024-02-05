package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"github.com/ahr-i/triton-agent/setting"
	"github.com/ahr-i/triton-agent/src/errController"
	"github.com/gorilla/mux"
)

type ClientRequest struct {
	Token    string `json:"token"`
	Response string `json:"response"`
}

func (h *Handler) inferHandler(w http.ResponseWriter, r *http.Request) {
	// Extract model information from the URL.
	vars := mux.Vars(r)
	model := vars["model"]
	version := vars["version"]

	query := r.URL.Query()
	address := query.Get("address")
	token := query.Get("key")

	// Extract the request from the body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errController.ErrorMessage(err)
		return
	}
	log.Println("* (System) Request: ▽▽▽▽▽▽▽▽▽▽")
	log.Println(string(body))

	response, err := requestToTriton(model, version, body)
	if err != nil {
		errController.ErrorMessage(err)
		return
	}

	err = sendToGateway(address, response, token)
	if err != nil {
		errController.ErrorMessage(err)
		return
	}

	rend.JSON(w, http.StatusOK, nil)
}

/* Send a request to the Triton server and return the response. */
func requestToTriton(model string, version string, request []byte) ([]byte, error) {
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

	return bodyBytes, nil
}

/* Pass the Triton's response to the Gateway. */
func sendToGateway(address string, response []byte, token string) error {
	// TCP connection
	senderConn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}
	defer senderConn.Close()

	// Message setting
	clientRequest := ClientRequest{
		Token:    token,
		Response: string(response),
	}
	jsonData, err := json.Marshal(clientRequest)
	if err != nil {
		return err
	}

	// Send Message
	_, err = senderConn.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}
