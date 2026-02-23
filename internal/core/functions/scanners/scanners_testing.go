package scanners

import "Scruticode/internal/core/types"

func UnitTestsConfigured(basePath string) types.BaseResponse {
	if directoryExists(basePath, "tests", "test", "unit") {
		return qualityCheckResponse(true)
	}

	return qualityCheckResponse(hasFileWithSuffix(basePath, "_test.go", ".spec.ts", ".spec.js", ".test.ts", ".test.js"))
}

func IntegrationTestsConfigured(basePath string) types.BaseResponse {
	if directoryExists(basePath, "integration", "integration-tests", "tests/integration") {
		return qualityCheckResponse(true)
	}

	return qualityCheckResponse(hasFileWithNameContaining(basePath, "integration", "it"))
}

func E2ETestsConfigured(basePath string) types.BaseResponse {
	if directoryExists(basePath, "e2e", "tests/e2e", "cypress", "playwright") {
		return qualityCheckResponse(true)
	}

	return qualityCheckResponse(hasFileWithNameContaining(basePath, "e2e", "playwright", "cypress"))
}

func CoverageConfigured(basePath string) types.BaseResponse {
	if anyFileExists(basePath, "codecov.yml", "codecov.yaml", ".coveragerc") {
		return qualityCheckResponse(true)
	}

	return qualityCheckResponse(hasScriptInPackageJSON(basePath, "coverage", "test:coverage"))
}

func StressTestsConfigured(basePath string) types.BaseResponse {
	if directoryExists(basePath, "stress", "load", "performance", "k6") {
		return qualityCheckResponse(true)
	}

	if anyFileExists(basePath, "k6.js", "k6.ts", "artillery.yml", "artillery.yaml") {
		return qualityCheckResponse(true)
	}

	return qualityCheckResponse(hasScriptInPackageJSON(basePath, "stress", "load", "performance"))
}
