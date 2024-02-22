package tritonController

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/ahr-i/triton-agent/modelStoreCommunicator"
	"github.com/ahr-i/triton-agent/setting"
)

func SetModel(provider string, model string, version string) error {
	filePath := fmt.Sprintf("%s/%s", setting.ModelsPath, provider)

	if err := makeFolder(filePath); err != nil {
		return err
	}

	modelFile, err := modelStoreCommunicator.GetModel(provider, model, version)
	if err != nil {
		return err
	}

	zipReader, err := zip.NewReader(bytes.NewReader(modelFile), int64(len(modelFile)))
	if err != nil {
		return err
	}

	for _, file := range zipReader.File {
		// 출력 경로 생성
		outputPath := filepath.Join(filePath, file.Name)

		// 디렉토리인 경우 생성
		if file.FileInfo().IsDir() {
			os.MkdirAll(outputPath, os.ModePerm)
			continue
		}

		// 파일 내용 읽기
		fileInZip, err := file.Open()
		if err != nil {
			panic(err)
		}

		// 경로에 해당하는 디렉토리 생성
		if err := os.MkdirAll(filepath.Dir(outputPath), os.ModePerm); err != nil {
			panic(err)
		}

		// 파일 생성 및 쓰기
		outputFile, err := os.Create(outputPath)
		if err != nil {
			panic(err)
		}
		if _, err := io.Copy(outputFile, fileInZip); err != nil {
			panic(err)
		}

		// 파일 닫기
		fileInZip.Close()
		outputFile.Close()
	}

	return nil
}
