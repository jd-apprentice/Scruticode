package scanners

import (
	"Scruticode/internal/core/types"
	"path/filepath"
)

func CIPipelineExists(basePath string) types.BaseResponse {
	if anyFileExists(basePath, ".gitlab-ci.yml", "azure-pipelines.yml", "bitbucket-pipelines.yml") {
		return qualityCheckResponse(true)
	}

	workflowsPath := filepath.Join(basePath, ".github", "workflows")

	return qualityCheckResponse(hasFileWithNameContaining(workflowsPath, "ci", "build", "test", "pipeline"))
}

func CDPipelineExists(basePath string) types.BaseResponse {
	if anyFileExists(basePath, "appspec.yml", "appspec.yaml", "skaffold.yaml", "skaffold.yml") {
		return qualityCheckResponse(true)
	}

	workflowsPath := filepath.Join(basePath, ".github", "workflows")

	return qualityCheckResponse(hasFileWithNameContaining(workflowsPath, "cd", "deploy", "release", "publish"))
}
