package tritonController

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ahr-i/triton-agent/setting"
)

func CheckModel(provider string, model string) bool {
	modelRepository := GetModelRepository()

	value, exist := modelRepository[provider][model]

	return value && exist
}

func PrintModelRepository() {
	modelRepository := GetModelRepository()

	for provider, models := range modelRepository {
		log.Println("Provider:", provider)
		log.Println("- Models:", models)
	}
}

func GetModelRepository() map[string]map[string]bool {
	return getDirectoriesWithDepth(setting.ModelsPath, 2)
}

func getDirectoriesWithDepth(basePath string, depth int) map[string]map[string]bool {
	result := make(map[string]map[string]bool)
	_ = filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			relativePath, _ := filepath.Rel(basePath, path)
			parts := strings.Split(relativePath, string(os.PathSeparator))

			if len(parts) == depth {
				parentDir := parts[0]
				currentDir := parts[1]

				if _, exists := result[parentDir]; !exists {
					result[parentDir] = make(map[string]bool)
				}
				result[parentDir][currentDir] = true
			}
		}

		return nil
	})

	return result
}
