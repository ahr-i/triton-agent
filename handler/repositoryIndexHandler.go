package handler

import (
	"net/http"

	"github.com/ahr-i/triton-agent/src/httpController"
	"github.com/ahr-i/triton-agent/src/logCtrlr"
	"github.com/ahr-i/triton-agent/tritonCommunicator"
)

func (h *Handler) repositoryIndexHandler(w http.ResponseWriter, r *http.Request) {
	modelInformation, err := tritonCommunicator.GetRepositoryIndex()
	if err != nil {
		logCtrlr.Error(err)
		rend.JSON(w, http.StatusBadRequest, nil)
		return
	}

	httpController.JSON(w, http.StatusOK, modelInformation)
}
