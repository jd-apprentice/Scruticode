package options

import (
	"Scruticode/src/constants"
	"Scruticode/src/functions/utils"
	"Scruticode/src/types"
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
			utils.LoggerErrorFile(constants.FileNotFound, folder)

			return types.BaseResponse{
				Status: constants.QualityCheckFailed,
			}
		}

		for _, file := range files {
			if strings.Contains(file.Name(), "Dockerfile") {
				utils.LoggerDebugFile(constants.FileFound, file.Name())

				return types.BaseResponse{
					Status: constants.QualityCheckSuccess,
				}
			}
		}
	}

	utils.LoggerErrorFile(constants.FileNotFound, strings.Join(PossibleFolderPath, ", "))

	return types.BaseResponse{
		Status: constants.QualityCheckFailed,
	}
}
