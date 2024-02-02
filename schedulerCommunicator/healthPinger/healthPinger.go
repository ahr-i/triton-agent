package healthPinger

import (
	"log"
	"time"

	"github.com/ahr-i/triton-agent/setting"
	"github.com/ahr-i/triton-agent/src/networkController"
)

var address string

func Enter() {
	log.Println("* (System) Send information to the Scheduler.")

	address = networkController.GetLocalIP() + ":" + setting.ServerPort

	alivePoster()
}

func alivePoster() {
	var cnt int = 0

	for {
		cnt++
		log.Printf("* (System) Send information to the Scheduler. (It is the %dth request)\n", cnt)

		postAlive()

		time.Sleep(5 * time.Second)
	}
}
