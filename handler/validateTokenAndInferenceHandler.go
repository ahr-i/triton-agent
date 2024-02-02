package handler

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ahr-i/triton-agent/schedulerCommunicator/taskTokenManager"
	"github.com/ahr-i/triton-agent/setting"
	"github.com/gorilla/mux"
)

func (h *Handler) validateTokenAndInferenceHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	token := query.Get("token")

	// Token Check
	exist := taskTokenManager.IsValidToken(token)
	if !exist {
		log.Println("** (ERROR) Invalid token value.")
		log.Println("** (ERROR) Error token:" + token)
		rend.JSON(w, http.StatusBadRequest, nil)

		return
	}

	log.Println("* (System) Use the token.")
	log.Println("* (System) Use token:" + token)
	taskTokenManager.UseToken(token)

	vars := mux.Vars(r)
	model := vars["model"]
	version := vars["version"]

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	// Triton Server Request Setting
	responseData := requestTriton(body, model, version)

	rend.JSON(w, http.StatusOK, string(responseData))
}

func requestTriton(request []byte, model string, version string) []byte {
	log.Println("* (System) Inference request: ▽▽▽▽▽▽▽▽▽▽")
	log.Println(string(request))

	url := "http://" + setting.TritonUrl + "/v2/models/" + model + "/versions/" + version + "/infer"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(request))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Triton Server Response
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return bodyBytes
}
