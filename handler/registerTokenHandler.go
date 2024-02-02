package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ahr-i/triton-agent/schedulerCommunicator/taskTokenManager"
)

type RequestData struct {
	Token string `json:"token"`
}

func (h *Handler) registerTokenHandler(w http.ResponseWriter, r *http.Request) {
	var requestData RequestData
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	log.Println("* (System) Register the token.")
	log.Println("* (System) Token:" + requestData.Token)
	taskTokenManager.Register(requestData.Token)

	rend.JSON(w, http.StatusOK, nil)
}
