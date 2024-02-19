package healthPinger

import (
	"log"
	"time"

	"github.com/ahr-i/triton-agent/setting"
)

var port string

func Enter() {
	port = setting.ServerPort

	alivePoster()
}

func alivePoster() {
	var cnt int = 0

	for {
		cnt++
		log.Printf("* (System) Send information to the Scheduler. (It is the %dth request)\n", cnt)

		postAlive()

		time.Sleep(8 * time.Second)
	}
}
