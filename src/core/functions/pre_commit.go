package functions

import (
	"Scruticode/src/core/types"
	"Scruticode/src/shared/constants"
	"log"
	"os"
)

func PreCommitExists(lang string, projectPath ...string) types.BaseResponse {
	basePath, err := getBasePath(projectPath...)
	if err != nil {
		log.Printf("%s: %s\n", constants.FileNotFound, err.Error())
		return types.BaseResponse{Status: constants.QualityCheckFailed}
	}

	if lang == "javascript" {
		packageJSONPath := basePath + string(os.PathSeparator) + "package.json"
		return checkHusky(basePath, packageJSONPath)
	}

	return checkPreCommitConfig(basePath)
}

func checkHusky(basePath, packageJSONPath string) types.BaseResponse {
	huskyPath := basePath + string(os.PathSeparator) + ".husky"
	if _, err := os.Stat(huskyPath); os.IsNotExist(err) {
		log.Printf("%s: .husky folder not found.\n", constants.CheckWarning)
		return types.BaseResponse{Status: constants.QualityCheckFailed}
	}
	log.Printf("%s: .husky folder found.\n", constants.CheckPassed)

	packageJSON, err := readAndParsePackageJSON(packageJSONPath)
	if err != nil {
		log.Printf("%s: %s\n", constants.FileNotFound, err.Error())
		return types.BaseResponse{Status: constants.QualityCheckFailed}
	}

	if !hasDependency(packageJSON, "husky") {
		log.Printf("%s: husky dependency not found in package.json.\n", constants.CheckWarning)
		return types.BaseResponse{Status: constants.QualityCheckFailed}
	}
	log.Printf("%s: husky dependency found in package.json.\n", constants.CheckPassed)

	return types.BaseResponse{Status: constants.QualityCheckSuccess}
}

func checkPreCommitConfig(basePath string) types.BaseResponse {
	yamlPath := basePath + string(os.PathSeparator) + ".pre-commit-config.yaml"
	ymlPath := basePath + string(os.PathSeparator) + ".pre-commit-config.yml"

	if _, err := os.Stat(yamlPath); err == nil {
		log.Printf("%s: .pre-commit-config.yaml found.\n", constants.CheckPassed)
		return types.BaseResponse{Status: constants.QualityCheckSuccess}
	}

	if _, err := os.Stat(ymlPath); err == nil {
		log.Printf("%s: .pre-commit-config.yml found.\n", constants.CheckPassed)
		return types.BaseResponse{Status: constants.QualityCheckSuccess}
	}

	log.Printf("%s: No .pre-commit-config file found.\n", constants.CheckWarning)
	return types.BaseResponse{Status: constants.QualityCheckFailed}
}
