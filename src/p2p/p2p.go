package p2p

import (
	"time"

	"github.com/ahr-i/triton-agent/setting"
	"github.com/ahr-i/triton-agent/src/networkController"
)

func Enter() {
	address := networkController.GetLocalIP() + ":" + setting.ServerPort

	alivePoster(address)
}

func alivePoster(address string) {
	for {
		postAlive(address)

		time.Sleep(5 * time.Second)
	}
}
