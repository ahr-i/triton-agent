package tritonController

import (
	"fmt"
	"log"
	"time"

	"github.com/anacrolix/torrent"
	"github.com/anacrolix/torrent/metainfo"
)

// 다운로드가 완료된 파일을 채널로 입력받아서 채널입력이 올때마다 시딩 리스트 추가해서 시딩
func Seeding(filepath string, channel *chan string) {
	port := 50000
	for {
		select {
		case fileName := <-*channel:
			go seedingAdditonal(filepath, fileName, port)
			port += 1
		}
	}
}

func seedingAdditonal(filepath string, fileName string, port int) {
	// 토렌트 클라이언트 설정
	cfg := torrent.NewDefaultClientConfig()
	cfg.Seed = true        // 시딩 활성화
	cfg.DataDir = filepath // 다운로드된 파일이 위치한 디렉토리
	cfg.ListenPort = port

	// 토렌트 클라이언트 생성
	cl, err := torrent.NewClient(cfg)
	if err != nil {
		log.Fatalf("error creating client: %s", err)
	}
	defer cl.Close()

	// 토렌트 파일 추가
	torrentPath := fmt.Sprintf("%s/%s", filepath, fileName) // 여기에 토렌트 파일 경로 입력
	metaInfo, err := metainfo.LoadFromFile(torrentPath)
	if err != nil {
		log.Fatalf("error loading torrent file: %s", err)
	}
	t, err := cl.AddTorrent(metaInfo)
	if err != nil {
		log.Fatalf("error adding torrent: %s", err)
	}

	// 시딩을 위해 클라이언트를 계속 실행
	log.Printf("Seeding %s", t.Name())

	go func() {
		for range time.Tick(10 * time.Second) {
			// 여기서 추가적으로 특정 조건을 확인하고 메시지를 출력할 수 있음
			peers := t.PeerConns()
			//log.Println(peers)
			if len(peers) != 0 {
				log.Printf("Downloading~. file: %s, peer: %d.\n", fileName, len(peers))
			}
		}
	}()

	select {}
}
