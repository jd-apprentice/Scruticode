package functions

import (
	scannerChecks "Scruticode/internal/core/functions/scanners"
	"Scruticode/internal/core/types"
	"Scruticode/internal/shared/constants"
	"Scruticode/internal/shared/utils"
	"fmt"
	"log"
	"os"
	"strings"
)

func ScanDirectory(path string, lang string, platform string) {
	log.Println("Scanning directory:", path)

	err := os.Chdir(path)
	if err != nil {
		log.Fatalf("Failed to change directory: %v", err)
	}

	configuration := ReadConfigFile()
	keyValues := ProcessConfigFile(configuration)
	results := RunScanners(keyValues, lang, platform)
	fmt.Println(FormatScanResults(results))
}

func RunScanners(keyValues []string, lang string, platform string) []types.CheckResult {
	var results []types.CheckResult

	actions := buildScannerActions()

	currentLang := lang
	currentPlatform := platform

	for _, pair := range keyValues {
		key, value := parseKeyValuePair(pair)
		if key == "" {
			continue
		}

		currentLang = updateLanguageConfig(key, value, currentLang)
		currentPlatform = updatePlatformConfig(key, value, currentPlatform)
		processLanguageOrPlatformConfig(key, value)

		if !isActionEnabled(value) {
			continue
		}

		results = append(results, executeScannerAction(actions, key, currentLang))
	}
	return results
}

func buildScannerActions() map[string]types.CheckFunc {
	return map[string]types.CheckFunc{
		"docker_compose": func(l string) types.CheckResult {
			return types.CheckResult{Name: "docker_compose", Passed: scannerChecks.DockerComposeExists(constants.CurrentPath).Status == constants.QualityCheckSuccess}
		},
		"dockerfile": func(l string) types.CheckResult {
			return types.CheckResult{Name: "dockerfile", Passed: scannerChecks.DockerfileExists(constants.CurrentPath).Status == constants.QualityCheckSuccess}
		},
		"readme": func(l string) types.CheckResult {
			return types.CheckResult{Name: "readme", Passed: scannerChecks.Readme(constants.ReadmeFilePath).Status == constants.QualityCheckSuccess}
		},
		"ci": func(l string) types.CheckResult {
			return types.CheckResult{Name: "ci", Passed: scannerChecks.CIPipelineExists(constants.CurrentPath).Status == constants.QualityCheckSuccess}
		},
		"cd": func(l string) types.CheckResult {
			return types.CheckResult{Name: "cd", Passed: scannerChecks.CDPipelineExists(constants.CurrentPath).Status == constants.QualityCheckSuccess}
		},
		"conventional_commits": func(l string) types.CheckResult {
			return types.CheckResult{Name: "conventional_commits", Passed: scannerChecks.ConventionalCommitsConfigured(constants.CurrentPath).Status == constants.QualityCheckSuccess}
		},
		"copilot": func(l string) types.CheckResult {
			return types.CheckResult{Name: "copilot", Passed: scannerChecks.CopilotRulesExists(constants.CopilotInstructionsPath).Status == constants.QualityCheckSuccess}
		},
		"cline": func(l string) types.CheckResult {
			return types.CheckResult{Name: "cline", Passed: scannerChecks.ClineRulesExists(constants.ClineRulesDirPath).Status == constants.QualityCheckSuccess}
		},
		"pre_commit": func(l string) types.CheckResult {
			return types.CheckResult{Name: "pre_commit", Passed: scannerChecks.PreCommitExists(l, constants.CurrentPath).Status == constants.QualityCheckSuccess}
		},
		"linter": func(l string) types.CheckResult {
			return types.CheckResult{Name: "linter", Passed: scannerChecks.LinterJavascriptExists(constants.CurrentPath).Status == constants.QualityCheckSuccess}
		},
		"formatter": func(l string) types.CheckResult {
			return types.CheckResult{Name: "formatter", Passed: scannerChecks.FormatterConfigured(constants.CurrentPath).Status == constants.QualityCheckSuccess}
		},
		"unit": func(l string) types.CheckResult {
			return types.CheckResult{Name: "unit", Passed: scannerChecks.UnitTestsConfigured(constants.CurrentPath).Status == constants.QualityCheckSuccess}
		},
		"integration": func(l string) types.CheckResult {
			return types.CheckResult{Name: "integration", Passed: scannerChecks.IntegrationTestsConfigured(constants.CurrentPath).Status == constants.QualityCheckSuccess}
		},
		"e2e": func(l string) types.CheckResult {
			return types.CheckResult{Name: "e2e", Passed: scannerChecks.E2ETestsConfigured(constants.CurrentPath).Status == constants.QualityCheckSuccess}
		},
		"coverage": func(l string) types.CheckResult {
			return types.CheckResult{Name: "coverage", Passed: scannerChecks.CoverageConfigured(constants.CurrentPath).Status == constants.QualityCheckSuccess}
		},
		"stress": func(l string) types.CheckResult {
			return types.CheckResult{Name: "stress", Passed: scannerChecks.StressTestsConfigured(constants.CurrentPath).Status == constants.QualityCheckSuccess}
		},
		"secrets": func(l string) types.CheckResult {
			return types.CheckResult{Name: "secrets", Passed: scannerChecks.SecretScanningConfigured(constants.CurrentPath).Status == constants.QualityCheckSuccess}
		},
		"iac": func(l string) types.CheckResult {
			return types.CheckResult{Name: "iac", Passed: scannerChecks.IACScanningConfigured(constants.CurrentPath).Status == constants.QualityCheckSuccess}
		},
		"code": func(l string) types.CheckResult {
			return types.CheckResult{Name: "code", Passed: scannerChecks.CodeSecurityScanningConfigured(constants.CurrentPath).Status == constants.QualityCheckSuccess}
		},
		"container": func(l string) types.CheckResult {
			return types.CheckResult{Name: "container", Passed: scannerChecks.ContainerSecurityScanningConfigured(constants.CurrentPath).Status == constants.QualityCheckSuccess}
		},
		"deps": func(l string) types.CheckResult {
			return types.CheckResult{Name: "deps", Passed: scannerChecks.DependencyScanningConfigured(constants.CurrentPath).Status == constants.QualityCheckSuccess}
		},
		"sast": func(l string) types.CheckResult {
			return types.CheckResult{Name: "sast", Passed: scannerChecks.SASTConfigured(constants.CurrentPath).Status == constants.QualityCheckSuccess}
		},
		"dast": func(l string) types.CheckResult {
			return types.CheckResult{Name: "dast", Passed: scannerChecks.DASTConfigured(constants.CurrentPath).Status == constants.QualityCheckSuccess}
		},
	}
}

func updateLanguageConfig(key, value, currentLang string) string {
	if key == "langs" && currentLang == "" {
		return utils.ToAbsoluteString(value)
	}
	return currentLang
}

func updatePlatformConfig(key, value, currentPlatform string) string {
	if key == "platforms" && currentPlatform == "" {
		return utils.ToAbsoluteString(value)
	}
	return currentPlatform
}

func processLanguageOrPlatformConfig(key, value string) {
	isLangOrPlatform := key == "langs" || key == "platforms"
	if isLangOrPlatform {
		keyAsString := utils.ToAbsoluteString(value)
		extraLangConfig(keyAsString)
		extraPlatformConfig(keyAsString)
	}
}

func executeScannerAction(actions map[string]types.CheckFunc, key, currentLang string) types.CheckResult {
	action, exists := actions[key]
	if !exists {
		log.Printf("No action found for key '%s'\n", key)
		return types.CheckResult{Name: key, Passed: false}
	}
	return action(currentLang)
}

func isActionEnabled(action string) bool {
	return strings.ToLower(action) == "true"
}

func parseKeyValuePair(pair string) (string, string) {
	parts := strings.SplitN(pair, "=", constants.IsKeyVal)
	if len(parts) != constants.IsKeyVal {
		return "", ""
	}

	key := strings.TrimSpace(parts[constants.IsFirstIndex])
	value := strings.TrimSpace(parts[constants.IsSecondIndex])

	return key, value
}

func FormatScanResults(results []types.CheckResult) string {
	var passedChecks []string
	var failedChecks []string
	var output strings.Builder

	for _, result := range results {
		if result.Passed {
			passedChecks = append(passedChecks, result.Name)
		} else {
			failedChecks = append(failedChecks, result.Name)
		}
	}

	output.WriteString(fmt.Sprintln("\n### Status Check"))
	output.WriteString(fmt.Sprintf("✅ Checks passed: %d\n", len(passedChecks)))
	output.WriteString(fmt.Sprintf("❌ Checks failed: %d\n", len(failedChecks)))

	output.WriteString("\n## 📚 List of checks passed\n")
	if len(passedChecks) == 0 {
		output.WriteString("- None\n")
	}

	if len(passedChecks) != 0 {
		for _, check := range passedChecks {
			output.WriteString(fmt.Sprintf("- %s\n", check))
		}
	}

	output.WriteString("\n## 📚 List of checks failed\n")
	if len(failedChecks) == 0 {
		output.WriteString("- None\n")
	}

	if len(failedChecks) != 0 {
		for _, check := range failedChecks {
			output.WriteString(fmt.Sprintf("- %s\n", check))
		}
	}

	return output.String()
}
