package functions

import (
	"Scruticode/internal/core/types"
	"Scruticode/internal/shared/constants"
	"log"
	"os"
)

func CopilotRulesExists(copilot_file string) types.BaseResponse {
	if _, err := os.Stat(copilot_file); err == nil {
		log.Printf("%s: copilot-instructions.md\n", constants.FileFound)
		return types.BaseResponse{Status: constants.QualityCheckSuccess}
	}

	return types.BaseResponse{Status: constants.QualityCheckFailed}
}
