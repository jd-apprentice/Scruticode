package scanners

import (
	"Scruticode/internal/core/types"
	"Scruticode/internal/shared/constants"
	"log"
	"os"
	"path/filepath"
)

const logFileStatusFormat = "%s: %s\n"

func ClineRulesExists(clineRulesPath string) types.BaseResponse {
	info, err := os.Stat(clineRulesPath)
	if err != nil || !info.IsDir() {
		return qualityCheckResponse(false)
	}

	entries, err := os.ReadDir(clineRulesPath)
	if err != nil || len(entries) == 0 {
		return qualityCheckResponse(false)
	}

	log.Printf(logFileStatusFormat, constants.FileFound, clineRulesPath)

	return qualityCheckResponse(true)
}

func CopilotRulesExists(copilotFile string) types.BaseResponse {
	if _, err := os.Stat(copilotFile); err == nil {
		log.Printf("%s: copilot-instructions.md\n", constants.FileFound)

		return qualityCheckResponse(true)
	}

	return qualityCheckResponse(false)
}

func DockerfileExists(basePath string) types.BaseResponse {
	filePath := filepath.Join(basePath, "Dockerfile")
	if _, err := os.Stat(filePath); err != nil {
		log.Printf(logFileStatusFormat, constants.FileNotFound, filePath)

		return qualityCheckResponse(false)
	}

	log.Printf(logFileStatusFormat, constants.FileFound, filePath)

	return qualityCheckResponse(true)
}

func Readme(path string) types.BaseResponse {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Printf(logFileStatusFormat, constants.FileNotFound, path)

		return qualityCheckResponse(false)
	}

	log.Printf(logFileStatusFormat, constants.FileFound, path)

	return qualityCheckResponse(true)
}
