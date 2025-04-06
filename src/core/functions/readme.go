package functions

import (
	"Scruticode/src/core/types"
	"Scruticode/src/shared/constants"
	"log"
	"os"
)

func Readme(path string) types.BaseResponse {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatalf("%s: %s\n", constants.FileNotFound, path)

		return types.BaseResponse{Status: constants.QualityCheckFailed}
	}

	log.Printf("%s: %s\n", constants.FileFound, path)

	return types.BaseResponse{Status: constants.QualityCheckSuccess}
}
