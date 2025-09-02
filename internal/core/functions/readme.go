package functions

import (
	"Scruticode/internal/core/types"
	"Scruticode/internal/shared/constants"
	"log"
	"os"
)

func Readme(path string) types.BaseResponse {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Printf("%s: %s\n", constants.FileNotFound, path)

		return types.BaseResponse{Status: constants.QualityCheckFailed}
	}

	log.Printf("%s: %s\n", constants.FileFound, path)

	return types.BaseResponse{Status: constants.QualityCheckSuccess}
}
