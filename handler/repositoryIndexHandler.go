package handler

import (
	"net/http"

	"github.com/ahr-i/triton-agent/tritonCommunicator"
)

func (h *Handler) repositoryIndexHandler(w http.ResponseWriter, r *http.Request) {
	modelInformation, err := tritonCommunicator.GetRepositoryIndex()
	if err != nil {
		panic(err)
	}

	rend.JSON(w, http.StatusOK, string(modelInformation))
}
