package functions

import (
	"Scruticode/internal/core/types"
	"Scruticode/internal/shared/constants"
	"log"
	"os"
	"strings"
)

func DockerfileExists(folder string) types.BaseResponse {
	const fatalMessage = "%s: %s\n"

	const dockerfile = "Dockerfile"

	files, err := os.ReadDir(folder)
	if err != nil {
		log.Fatalf(fatalMessage, constants.FileNotFound, folder)
	}

	for _, file := range files {
		if strings.Contains(file.Name(), dockerfile) {
			log.Printf(fatalMessage, constants.FileFound, file.Name())

			return types.BaseResponse{
				Status: constants.QualityCheckSuccess,
			}
		}
	}

	log.Fatalf(fatalMessage, constants.FileNotFound, folder+"/"+dockerfile)

	return types.BaseResponse{
		Status: constants.QualityCheckFailed,
	}
}
