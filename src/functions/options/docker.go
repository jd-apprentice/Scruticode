package options

import (
	"Scruticode/src/constants"
	"Scruticode/src/functions/utils"
	"Scruticode/src/types"
	"os"
	"strings"
)

func Dockerfile() types.BaseResponse {
	const DockerFilePath = "Dockerfile"
	PossibleFolderPath := []string{
		".",
		"docker",
		"infra",
	}

	for _, folder := range PossibleFolderPath {
		path := folder + "/" + DockerFilePath
		if _, err := os.Stat(path); err == nil {
			utils.LoggerDebugFile(constants.FileFound, path)

			return types.BaseResponse{
				Status: constants.QualityCheckSuccess,
			}
		}
	}

	utils.LoggerErrorFile(constants.FileNotFound, strings.Join(PossibleFolderPath, ", "))

	return types.BaseResponse{
		Status: constants.QualityCheckFailed,
	}
}
