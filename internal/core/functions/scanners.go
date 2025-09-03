package functions

import (
	"Scruticode/internal/core/types"
	"Scruticode/internal/shared/constants"
	"Scruticode/internal/shared/utils"
	"fmt"
	"log"
	"os"
	"strings"
)

func ScanDirectory(path string) {
	log.Println("Scanning directory:", path)

	err := os.Chdir(path)
	if err != nil {
		log.Fatalf("Failed to change directory: %v", err)
	}

	configuration := ReadConfigFile()
	keyValues := ProcessConfigFile(configuration)
	results := RunScanners(keyValues)
	fmt.Println(FormatScanResults(results))
}

func RunScanners(keyValues []string) []types.CheckResult {
	var lang string
	var results []types.CheckResult

	var actions = map[string]types.CheckFunc{
		"docker_compose": func(l string) types.CheckResult { return types.CheckResult{Name: "docker_compose", Passed: false} },
		"dockerfile": func(l string) types.CheckResult {
			return types.CheckResult{Name: "dockerfile", Passed: DockerfileExists(constants.CurrentPath).Status == constants.QualityCheckSuccess}
		},
		"readme": func(l string) types.CheckResult {
			return types.CheckResult{Name: "readme", Passed: Readme(constants.ReadmeFilePath).Status == constants.QualityCheckSuccess}
		},
		"ci": func(l string) types.CheckResult { return types.CheckResult{Name: "ci", Passed: false} },
		"cd": func(l string) types.CheckResult { return types.CheckResult{Name: "cd", Passed: false} },
		"conventional_commits": func(l string) types.CheckResult {
			return types.CheckResult{Name: "conventional_commits", Passed: false}
		},
		"copilot": func(l string) types.CheckResult {
			return types.CheckResult{Name: "copilot", Passed: CopilotRulesExists(constants.CopilotInstructionsPath).Status == constants.QualityCheckSuccess}
		},
		"pre_commit": func(l string) types.CheckResult {
			return types.CheckResult{Name: "pre_commit", Passed: PreCommitExists(l, constants.CurrentPath).Status == constants.QualityCheckSuccess}
		},
		"linter": func(l string) types.CheckResult {
			return types.CheckResult{Name: "linter", Passed: LinterJavascriptExists(constants.CurrentPath).Status == constants.QualityCheckSuccess}
		},
		"formatter":   func(l string) types.CheckResult { return types.CheckResult{Name: "formatter", Passed: false} },
		"unit":        func(l string) types.CheckResult { return types.CheckResult{Name: "unit", Passed: false} },
		"integration": func(l string) types.CheckResult { return types.CheckResult{Name: "integration", Passed: false} },
		"e2e":         func(l string) types.CheckResult { return types.CheckResult{Name: "e2e", Passed: false} },
		"coverage":    func(l string) types.CheckResult { return types.CheckResult{Name: "coverage", Passed: false} },
		"stress":      func(l string) types.CheckResult { return types.CheckResult{Name: "stress", Passed: false} },
		"secrets":     func(l string) types.CheckResult { return types.CheckResult{Name: "secrets", Passed: false} },
		"iac":         func(l string) types.CheckResult { return types.CheckResult{Name: "iac", Passed: false} },
		"code":        func(l string) types.CheckResult { return types.CheckResult{Name: "code", Passed: false} },
		"container":   func(l string) types.CheckResult { return types.CheckResult{Name: "container", Passed: false} },
		"deps":        func(l string) types.CheckResult { return types.CheckResult{Name: "deps", Passed: false} },
		"sast":        func(l string) types.CheckResult { return types.CheckResult{Name: "sast", Passed: false} },
		"dast":        func(l string) types.CheckResult { return types.CheckResult{Name: "dast", Passed: false} },
	}

	const emptyAsString = ""
	for _, pair := range keyValues {
		key, value := parseKeyValuePair(pair)
		if key == emptyAsString {
			continue
		}

		keyAsString := utils.ToAbsoluteString(value)
		if key == "langs" {
			lang = keyAsString
		}
		keyLangOrPlatform := key == "langs" || key == "platforms"

		if keyLangOrPlatform {
			extraLangConfig(keyAsString)
			extraPlatformConfig(keyAsString)
		}

		if !isActionEnabled(value) {
			continue
		}

		action, exists := actions[key]
		if !exists {
			log.Printf("No action found for key '%s'\n", key)
			results = append(results, types.CheckResult{Name: key, Passed: false})
			continue
		}
		results = append(results, action(lang))
	}
	return results
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
