package functions

import (
	"Scruticode/internal/core/types"
	"Scruticode/internal/shared/constants"
	"log"
	"os"
)

func CopilotRulesExists(copilotFile string) types.BaseResponse {
	if _, err := os.Stat(copilotFile); err == nil {
		log.Printf("%s: copilot-instructions.md\n", constants.FileFound)
		return types.BaseResponse{Status: constants.QualityCheckSuccess}
	}

	return types.BaseResponse{Status: constants.QualityCheckFailed}
}
