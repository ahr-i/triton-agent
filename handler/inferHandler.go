package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"runtime"

	"github.com/ahr-i/triton-agent/setting"
	"github.com/ahr-i/triton-agent/src/errController"
	"github.com/gorilla/mux"
)

type ClientRequest struct {
	Key      string `json:"key"`
	Response string `json:"response"`
}

type TritonOutput struct {
	Outputs []struct {
		Data []float32 `json:"data"`
	} `json:"outputs"`
}

func (h *Handler) InferHandler(w http.ResponseWriter, r *http.Request) {
	_, fp, _, _ := runtime.Caller(1)

	vars := mux.Vars(r)
	model := vars["model"]
	version := vars["version"]

	query := r.URL.Query()
	address := query.Get("address")
	key := query.Get("key")

	url := "http://" + setting.TritonUrl + "/v2/models/" + model + "/versions/" + version + "/infer"

	/* */
	body, err := ioutil.ReadAll(r.Body)

	log.Println(string(body))
	/* */

	req, err_ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	errController.ErrorCheck(err_, "HTTP REQUEST ERROR", fp)
	req.Header.Set("Content-Type", "application/json")

	// Triton Server Response
	client := &http.Client{}
	resp, err_ := client.Do(req)
	errController.ErrorCheck(err_, "HTTP RESPONSE ERROR", fp)
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	clientRequest := ClientRequest{
		Key:      key,
		Response: string(bodyBytes),
	}

	// Request
	senderConn, err := net.Dial("tcp", address)
	if err != nil {
		panic(err)
	}
	defer senderConn.Close()

	jsonData, err := json.Marshal(clientRequest)
	if err != nil {
		panic(err)
	}

	_, err = senderConn.Write(jsonData)
	if err != nil {
		panic(err)
	}

	rend.JSON(w, http.StatusOK, nil)
}
