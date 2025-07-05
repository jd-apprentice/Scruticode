package functions

import (
	"Scruticode/src/core/types"
	"Scruticode/src/shared/constants"
	"log"
	"os"
	"strings"
)

func DockerfileExists(folder string) types.BaseResponse {
	const fatalMessage = "%s: %s\n"

	files, err := os.ReadDir(folder)
	if err != nil {
		log.Fatalf(fatalMessage, constants.FileNotFound, folder)
	}

	for _, file := range files {
		if strings.Contains(file.Name(), "Dockerfile") {
			log.Printf(fatalMessage, constants.FileFound, file.Name())

			return types.BaseResponse{
				Status: constants.QualityCheckSuccess,
			}
		}
	}

	log.Fatalf(fatalMessage, constants.FileNotFound, folder)

	return types.BaseResponse{
		Status: constants.QualityCheckFailed,
	}
}
