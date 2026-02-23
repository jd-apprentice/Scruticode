package scanners

import (
	"Scruticode/internal/core/types"
	"Scruticode/internal/shared/constants"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func qualityCheckResponse(passed bool) types.BaseResponse {
	if passed {
		return types.BaseResponse{Status: constants.QualityCheckSuccess}
	}

	return types.BaseResponse{Status: constants.QualityCheckFailed}
}

func anyFileExists(basePath string, relativePaths ...string) bool {
	for _, relativePath := range relativePaths {
		filePath := filepath.Join(basePath, relativePath)
		if _, err := os.Stat(filePath); err == nil {
			return true
		}
	}

	return false
}

func directoryExists(basePath string, relativePaths ...string) bool {
	for _, relativePath := range relativePaths {
		dirPath := filepath.Join(basePath, relativePath)
		info, err := os.Stat(dirPath)
		if err != nil {
			continue
		}

		if info.IsDir() {
			return true
		}
	}

	return false
}

func hasFileWithSuffix(basePath string, suffixes ...string) bool {
	found := false
	err := filepath.WalkDir(basePath, func(path string, d os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return nil
		}

		if d.IsDir() {
			return nil
		}

		for _, suffix := range suffixes {
			if strings.HasSuffix(strings.ToLower(path), strings.ToLower(suffix)) {
				found = true
				return filepath.SkipAll
			}
		}

		return nil
	})

	if err != nil {
		return false
	}

	return found
}

func hasFileWithNameContaining(basePath string, fragments ...string) bool {
	entries, err := os.ReadDir(basePath)
	if err != nil {
		return false
	}

	for _, entry := range entries {
		name := strings.ToLower(entry.Name())
		for _, fragment := range fragments {
			if strings.Contains(name, strings.ToLower(fragment)) {
				return true
			}
		}
	}

	return false
}

func hasScriptInPackageJSON(basePath string, scripts ...string) bool {
	packageJSONPath := filepath.Join(basePath, "package.json")
	packageJSON, err := readAndParsePackageJSON(packageJSONPath)
	if err != nil {
		return false
	}

	for _, script := range scripts {
		if hasScript(packageJSON, script) {
			return true
		}
	}

	return false
}

func readAndParsePackageJSON(path string) (map[string]interface{}, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var packageJSON map[string]interface{}
	if err := json.Unmarshal(content, &packageJSON); err != nil {
		return nil, err
	}

	return packageJSON, nil
}

func hasScript(packageJSON map[string]interface{}, script string) bool {
	scripts, found := packageJSON["scripts"].(map[string]interface{})
	if !found {
		return false
	}

	_, exists := scripts[script]

	return exists
}

func hasDependency(packageJSON map[string]interface{}, dependency string) bool {
	if dependencies, ok := packageJSON["dependencies"].(map[string]interface{}); ok {
		if _, exists := dependencies[dependency]; exists {
			return true
		}
	}

	if devDependencies, ok := packageJSON["devDependencies"].(map[string]interface{}); ok {
		if _, exists := devDependencies[dependency]; exists {
			return true
		}
	}

	return false
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
