package scanners

import (
	"Scruticode/internal/core/types"
	"path/filepath"
)

const (
	githubDirName      = ".github"
	workflowsDirName   = "workflows"
	dependabotFilePath = ".github/dependabot.yml"
)

func SecretScanningConfigured(basePath string) types.BaseResponse {
	if anyFileExists(basePath,
		".gitleaks.toml",
		".gitleaks.yml",
		".gitleaks.yaml",
		".trufflehog.yml",
		".trufflehog.yaml",
	) {
		return qualityCheckResponse(true)
	}

	workflowsPath := filepath.Join(basePath, githubDirName, workflowsDirName)

	return qualityCheckResponse(hasFileWithNameContaining(workflowsPath, "secret", "gitleaks", "trufflehog"))
}

func IACScanningConfigured(basePath string) types.BaseResponse {
	if hasFileWithSuffix(basePath, ".tf", ".tfvars") {
		return qualityCheckResponse(true)
	}

	workflowsPath := filepath.Join(basePath, githubDirName, workflowsDirName)

	return qualityCheckResponse(hasFileWithNameContaining(workflowsPath, "iac", "terraform", "checkov", "tfsec"))
}

func CodeSecurityScanningConfigured(basePath string) types.BaseResponse {
	if anyFileExists(basePath, "sonar-project.properties", ".semgrep.yml", ".semgrep.yaml") {
		return qualityCheckResponse(true)
	}

	workflowsPath := filepath.Join(basePath, githubDirName, workflowsDirName)

	return qualityCheckResponse(hasFileWithNameContaining(workflowsPath, "codeql", "sonar", "semgrep"))
}

func DependencyScanningConfigured(basePath string) types.BaseResponse {
	if anyFileExists(basePath, "renovate.json", ".renovaterc", "dependabot.yml", dependabotFilePath) {
		return qualityCheckResponse(true)
	}

	workflowsPath := filepath.Join(basePath, githubDirName, workflowsDirName)

	return qualityCheckResponse(hasFileWithNameContaining(workflowsPath, "dependabot", "dependency", "deps"))
}

func SASTConfigured(basePath string) types.BaseResponse {
	workflowsPath := filepath.Join(basePath, githubDirName, workflowsDirName)

	return qualityCheckResponse(hasFileWithNameContaining(workflowsPath, "sast", "codeql", "semgrep"))
}

func DASTConfigured(basePath string) types.BaseResponse {
	workflowsPath := filepath.Join(basePath, githubDirName, workflowsDirName)

	return qualityCheckResponse(hasFileWithNameContaining(workflowsPath, "dast", "zap", "nikto"))
}
