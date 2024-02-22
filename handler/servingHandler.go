package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ahr-i/triton-agent/src/logCtrlr"
	"github.com/ahr-i/triton-agent/tritonController"
)

type servingInformation struct {
	Provider string `json:"id"`
	FileName string `json:"filename"`
	Address  string `json:"addr"`
}

func (h *Handler) servingHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logCtrlr.Error(err)
		return
	}
	defer r.Body.Close()

	var response servingInformation
	if err := json.Unmarshal(body, &response); err != nil {
		logCtrlr.Error(err)
		return
	}

	tritonController.SetModel(response.Provider, response.FileName, "1")

	rend.JSON(w, http.StatusOK, nil)
}
