package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/ahr-i/triton-agent/setting"
	"github.com/ahr-i/triton-agent/src/logCtrlr"
	"github.com/ahr-i/triton-agent/tritonController"
)

type servingInformation struct {
	Provider  string `json:"id"`
	ModelName string `json:"model_name"`
	Version   string `json:"version"`
	FileName  string `json:"filename"`
	Address   string `json:"addr"`
}

/* Downloading the model upon request. */
func (h *Handler) servingHandler(w http.ResponseWriter, r *http.Request, channel *chan string) {
	// Reading the request body.
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logCtrlr.Error(err)
		rend.JSON(w, http.StatusBadRequest, nil)
		return
	}
	defer r.Body.Close()

	// Performing JSON parsing.
	var response servingInformation
	if err := json.Unmarshal(body, &response); err != nil {
		logCtrlr.Error(err)
		rend.JSON(w, http.StatusBadRequest, nil)
		return
	}
	logCtrlr.Log("Request: ▽▽▽▽▽▽▽▽▽▽")
	log.Println("Provider:", response.Provider)
	log.Println("ModelName:", response.ModelName)
	log.Println("Version:", response.Version)
	log.Println("File name:", response.FileName)
	log.Println("ModelStore Address:", response.Address)

	setting.ModelStoreUrl = response.Address

	// Initiating model download and folder setup according to the request.
	if err := tritonController.SetModel(response.Provider, response.ModelName, response.Version, response.FileName, channel); err != nil {
		logCtrlr.Error(err)
		rend.JSON(w, http.StatusBadRequest, nil)
		return
	}

	rend.JSON(w, http.StatusOK, nil)
}
