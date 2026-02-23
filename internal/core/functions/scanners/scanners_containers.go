package scanners

import (
	"Scruticode/internal/core/types"
	"Scruticode/internal/shared/constants"
	"path/filepath"
)

func DockerComposeExists(basePath string) types.BaseResponse {
	return qualityCheckResponse(anyFileExists(basePath,
		"docker-compose.yml",
		"docker-compose.yaml",
		"compose.yml",
		"compose.yaml",
	))
}

func ContainerSecurityScanningConfigured(basePath string) types.BaseResponse {
	dockerComposeMissing := DockerComposeExists(basePath).Status != constants.QualityCheckSuccess
	if dockerComposeMissing && !anyFileExists(basePath, "Dockerfile") {
		return qualityCheckResponse(false)
	}

	workflowsPath := filepath.Join(basePath, ".github", "workflows")

	return qualityCheckResponse(hasFileWithNameContaining(workflowsPath, "trivy", "grype", "container"))
}
