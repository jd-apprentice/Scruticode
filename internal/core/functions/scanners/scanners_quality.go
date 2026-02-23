package scanners

import (
	"Scruticode/internal/core/types"
	"path/filepath"
)

func ConventionalCommitsConfigured(basePath string) types.BaseResponse {
	if anyFileExists(basePath,
		".commitlintrc",
		".commitlintrc.json",
		".commitlintrc.yml",
		".commitlintrc.yaml",
		"commitlint.config.js",
		"commitlint.config.cjs",
	) {
		return qualityCheckResponse(true)
	}

	workflowsPath := filepath.Join(basePath, ".github", "workflows")

	return qualityCheckResponse(hasFileWithNameContaining(workflowsPath, "commitlint", "conventional"))
}

func FormatterConfigured(basePath string) types.BaseResponse {
	if anyFileExists(basePath,
		".prettierrc",
		".prettierrc.json",
		".prettierrc.yml",
		".prettierrc.yaml",
		".clang-format",
		".editorconfig",
		".golangci.yaml",
	) {
		return qualityCheckResponse(true)
	}

	return qualityCheckResponse(hasScriptInPackageJSON(basePath, "format", "fmt"))
}
