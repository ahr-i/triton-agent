package tritonController

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/ahr-i/triton-agent/modelStoreCommunicator"
	"github.com/ahr-i/triton-agent/schedulerCommunicator/healthPinger"
	"github.com/ahr-i/triton-agent/setting"
	"github.com/ahr-i/triton-agent/src/logCtrlr"
)

func SetModel(provider string, model string, version string, filename string) error {
	filePath := fmt.Sprintf("%s/%s", setting.ModelsPath, provider)

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

	// Unzipping the file.
	logCtrlr.Log("Unzip the model.")
	zipReader, err := zip.NewReader(bytes.NewReader(modelFile), int64(len(modelFile)))
	if err != nil {
		return err
	}

	// Saving the model to the specified path.
	log.Println("Saving the model.")
	for _, file := range zipReader.File {
		// Creating the output path.
		outputPath := filepath.Join(filePath, file.Name)

		// If it is a directory, create it.
		if file.FileInfo().IsDir() {
			os.MkdirAll(outputPath, os.ModePerm)
			continue
		}

		// Read file contents.
		fileInZip, err := file.Open()
		if err != nil {
			return err
		}

		// Create a directory corresponding to the path.
		if err := os.MkdirAll(filepath.Dir(outputPath), os.ModePerm); err != nil {
			return err
		}

		// Create and write to a file.
		outputFile, err := os.Create(outputPath)
		if err != nil {
			return err
		}
		if _, err := io.Copy(outputFile, fileInZip); err != nil {
			return err
		}

		// Close the file.
		fileInZip.Close()
		outputFile.Close()
	}

	healthPinger.UpdateModel(provider, model, version)
	return nil
}
