package options

import (
	"Scruticode/src/constants"
	"Scruticode/src/functions/utils"
	"Scruticode/src/types"
	"os"
)

func Readme() types.BaseResponse {
	// LOGIC
	if _, err := os.Stat(constants.ReadmeFilePath); os.IsNotExist(err) {
		utils.LoggerErrorFile(constants.FileNotFound, constants.ReadmeFilePath)

		return types.BaseResponse{
			Status: constants.QualityCheckFailed,
		}
	}

	// LOG
	utils.LoggerDebugFile(constants.FileFound, constants.ReadmeFilePath)

	// RESPONSE
	return types.BaseResponse{
		Status: constants.QualityCheckSuccess,
	}
}
