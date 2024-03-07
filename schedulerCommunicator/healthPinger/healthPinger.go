package healthPinger

import (
	"log"
	"net"
	"os/exec"
	"strings"
	"time"

	"github.com/ahr-i/triton-agent/setting"
)

var port string
var gpuName string

// var model_info map[string]map[string]TaskInfo = make(map[string]map[string]TaskInfo)
var model_info map[string]map[string]TaskInfo

type TaskInfo struct {
	LoadedAmount         int     `json:"loaded_amount"`
	AverageInferenceTime float32 `json:"average_inference_time"`
}

func Enter() {
	model_info = make(map[string]map[string]TaskInfo)

	port = setting.ServerPort
	cmd := exec.Command("nvidia-smi", "--query-gpu=name", "--format=csv,noheader")
	output, err := cmd.Output()
	if err != nil {
		log.Fatal("failed to get GPU name: ", err)
	}

	gpuName = strings.TrimSpace(string(output))

	alivePoster()
}

func alivePoster() {
	var cnt int = 0

	log.Println("tcp 오픈")
	ln, err := net.Listen("tcp", ":6934")
	if err != nil {
		log.Fatal("헬스체커용 tcp 오픈 실패", err)
	}

	go func() {
		log.Println("승인 중")
		_, err = ln.Accept()
		if err != nil {
			log.Fatal("Acppea", err)
		}

		log.Println("헬스체킹용 tcp 연결 성공")
	}()

	for {
		cnt++
		log.Printf("* (System) Send information to the Scheduler. (It is the %dth request)\n", cnt)

		postAlive()

		time.Sleep(8 * time.Second)
	}
}
