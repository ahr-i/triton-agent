package tritonController

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/ahr-i/triton-agent/modelStoreCommunicator"
	"github.com/ahr-i/triton-agent/schedulerCommunicator/healthPinger"
	"github.com/ahr-i/triton-agent/setting"
	"github.com/ahr-i/triton-agent/src/logCtrlr"
	"github.com/anacrolix/torrent"
	"github.com/anacrolix/torrent/metainfo"
)

func SetModel(provider string, model string, version string, filename string, channel *chan string) error {
	filePath := fmt.Sprintf("%s/%s", setting.ModelsPath, provider)
	fileName := fmt.Sprintf("%s@%s<%s>.torrent", provider, model, version)

	// Creating the provider folder.
	// If the provider folder already exists, it will not be created.
	logCtrlr.Log("Create provider folder.")
	if err := makeFolder(filePath); err != nil {
		return err
	}

	// downloading the model from the Model Store.
	logCtrlr.Log("Request model download to the Model Store.")
	modelFile, err := modelStoreCommunicator.GetModel(provider, model, version, filename)
	if err != nil {
		return err
	}
	logCtrlr.Log("Successfully completed the model download.")

	// Create a directory corresponding to the path.
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}

	// Create and write to a file.
	file, err := os.Create(filePath + fileName)
	if err != nil {
		return err
	}
	file.Write(modelFile)

	// File Download by Torrent
	// 토렌트 클라이언트 설정
	cfg := torrent.NewDefaultClientConfig()
	cfg.Seed = true // 시딩 활성화

	// 토렌트 클라이언트 생성
	cl, err := torrent.NewClient(cfg)
	if err != nil {
		log.Fatalf("error creating client: %s", err)
	}
	defer cl.Close()

	// 토렌트 파일 추가
	torrentPath := filePath + "/" + fileName // 여기에 토렌트 파일 경로 입력
	metaInfo, err := metainfo.LoadFromFile(torrentPath)
	if err != nil {
		log.Fatalf("error loading torrent file: %s", err)
	}
	t, err := cl.AddTorrent(metaInfo)
	if err != nil {
		log.Fatalf("error adding torrent: %s", err)
	}

	<-t.GotInfo() // 토렌트 정보를 받을 때까지 대기
	t.DownloadAll()

	// 파일 다운로드 대기
	if cl.WaitAll() {

		log.Printf("Downloaded %s", t.Name())
		*channel <- fileName

		healthPinger.UpdateModel(provider, model, version)
	}
	return nil
}
