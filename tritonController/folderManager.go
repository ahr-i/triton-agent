package tritonController

import (
	"log"
	"os"
)

func makeFolder(path string) error {
	log.Println("* (SYSTEM) Create folder:", path)
	if err := os.MkdirAll(path, 0755); err != nil {
		return err
	}

	return nil
}
