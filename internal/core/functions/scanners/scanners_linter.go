package scanners

import (
	"Scruticode/internal/core/types"
	"Scruticode/internal/shared/constants"
	"log"
	"os"
	"path/filepath"
)

const logErrorFormat = "%s: %s\n"

func LinterJavascriptExists(projectPath ...string) types.BaseResponse {
	const (
		eslintDependency = "eslint"
		lintScript       = "lint"
	)

	basePath, err := getBasePath(projectPath...)
	if err != nil {
		log.Printf(logErrorFormat, constants.FileNotFound, err.Error())

		return qualityCheckResponse(false)
	}

	packageJSONPath := filepath.Join(basePath, "package.json")
	packageJSON, err := readAndParsePackageJSON(packageJSONPath)
	if err != nil {
		log.Printf(logErrorFormat, constants.FileNotFound, err.Error())

		return qualityCheckResponse(false)
	}

	if !hasDependency(packageJSON, eslintDependency) {
		log.Printf("%s: %s not found in dependencies.\n", constants.CheckWarning, eslintDependency)

		return qualityCheckResponse(false)
	}

	if !hasScript(packageJSON, lintScript) {
		log.Printf("%s: %s script not found in package.json.\n", constants.CheckWarning, lintScript)

		return qualityCheckResponse(false)
	}

	return qualityCheckResponse(true)
}

func checkHusky(basePath string) types.BaseResponse {
	huskyPath := filepath.Join(basePath, ".husky")
	if _, err := os.Stat(huskyPath); os.IsNotExist(err) {
		log.Printf("%s: .husky folder not found.\n", constants.CheckWarning)

		return qualityCheckResponse(false)
	}

	packageJSONPath := filepath.Join(basePath, "package.json")
	packageJSON, err := readAndParsePackageJSON(packageJSONPath)
	if err != nil {
		log.Printf(logErrorFormat, constants.FileNotFound, err.Error())

		return qualityCheckResponse(false)
	}

	if !hasDependency(packageJSON, "husky") {
		log.Printf("%s: husky dependency not found in package.json.\n", constants.CheckWarning)

		return qualityCheckResponse(false)
	}

	return qualityCheckResponse(true)
}

func checkPreCommitConfig(basePath string) types.BaseResponse {
	if anyFileExists(basePath, ".pre-commit-config.yaml", ".pre-commit-config.yml") {
		return qualityCheckResponse(true)
	}

	return qualityCheckResponse(false)
}

func PreCommitExists(lang string, projectPath ...string) types.BaseResponse {
	basePath, err := getBasePath(projectPath...)
	if err != nil {
		log.Printf(logErrorFormat, constants.FileNotFound, err.Error())

		return qualityCheckResponse(false)
	}

	if lang == "javascript" {
		return checkHusky(basePath)
	}

	return checkPreCommitConfig(basePath)
}
