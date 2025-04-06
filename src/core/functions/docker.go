package functions

import (
	"Scruticode/src/core/types"
	"Scruticode/src/shared/constants"
	"log"
	"os"
	"strings"
)

func DockerfileExists() types.BaseResponse {
	PossibleFolderPath := []string{
		".",
		"docker",
		"infra",
	}

	for _, folder := range PossibleFolderPath {
		files, err := os.ReadDir(folder)
		if err != nil {
			log.Fatalf("%s: %s\n", constants.FileNotFound, folder)

			return types.BaseResponse{
				Status: constants.QualityCheckFailed,
			}
		}

		for _, file := range files {
			if strings.Contains(file.Name(), "Dockerfile") {
				log.Printf("%s: %s\n", constants.FileFound, file.Name())

				return types.BaseResponse{
					Status: constants.QualityCheckSuccess,
				}
			}
		}
	}

	log.Fatalf("%s: %s\n", constants.FileNotFound, strings.Join(PossibleFolderPath, ", "))

	return types.BaseResponse{
		Status: constants.QualityCheckFailed,
	}
}
