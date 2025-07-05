package functions

import (
	"Scruticode/src/core/types"
	"Scruticode/src/shared/constants"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func LinterJavascriptExists(projectPath ...string) types.BaseResponse {
	const (
		fatalMessage     = "%s: %s\n"
		packageJSONFile  = "package.json"
		eslintDependency = "eslint"
		lintScript       = "lint"
	)

	basePath, err := getBasePath(projectPath...)
	if err != nil {
		log.Printf(fatalMessage, constants.FileNotFound, err.Error())

		return types.BaseResponse{Status: constants.QualityCheckFailed}
	}

	packageJSONPath := basePath + string(os.PathSeparator) + packageJSONFile

	packageJSON, err := readAndParsePackageJSON(packageJSONPath)
	if err != nil {
		log.Printf(fatalMessage, constants.FileNotFound, err.Error())

		return types.BaseResponse{Status: constants.QualityCheckFailed}
	}

	if !hasDependency(packageJSON, eslintDependency) {
		log.Printf("%s: %s not found in dependencies.\n", constants.CheckWarning, eslintDependency)

		return types.BaseResponse{Status: constants.QualityCheckFailed}
	}
	log.Printf("%s: %s found in dependencies.\n", constants.CheckPassed, eslintDependency)

	if !hasScript(packageJSON, lintScript) {
		log.Printf("%s: %s script not found in package.json.\n", constants.CheckWarning, lintScript)

		return types.BaseResponse{Status: constants.QualityCheckFailed}
	}
	log.Printf("%s: %s script found in package.json.\n", constants.CheckPassed, lintScript)

	return types.BaseResponse{Status: constants.QualityCheckSuccess}
}

func getBasePath(projectPath ...string) (string, error) {
	if len(projectPath) > 0 && projectPath[0] != "" {
		return projectPath[0], nil
	}
	currentDirectory, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("error getting current directory: %w", err)
	}

	return currentDirectory, nil
}

func readAndParsePackageJSON(path string) (map[string]interface{}, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading package.json: %w", err)
	}
	var packageJSON map[string]interface{}
	if err := json.Unmarshal(content, &packageJSON); err != nil {
		return nil, fmt.Errorf("error parsing package.json: %w", err)
	}

	return packageJSON, nil
}

func hasDependency(packageJSON map[string]interface{}, dependency string) bool {
	if deps, ok := packageJSON["dependencies"].(map[string]interface{}); ok {
		if _, exists := deps[dependency]; exists {
			return true
		}
	}
	if devDeps, ok := packageJSON["devDependencies"].(map[string]interface{}); ok {
		if _, exists := devDeps[dependency]; exists {
			return true
		}
	}
	return false
}

func hasScript(packageJSON map[string]interface{}, script string) bool {
	scripts, found := packageJSON["scripts"].(map[string]interface{})
	if !found {
		return false
	}
	_, exists := scripts[script]

	return exists
}
