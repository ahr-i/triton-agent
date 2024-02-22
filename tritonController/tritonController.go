package tritonController

import (
	"log"

	"github.com/ahr-i/triton-agent/src/logCtrlr"
)

var modelRepository string

func Init(repository string) {
	logCtrlr.Log("Set up the model repository.")
	log.Println("Model repository:", repository)

	modelRepository = repository
}
