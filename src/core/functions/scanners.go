package functions

import (
	"Scruticode/src/shared/constants"
	"Scruticode/src/shared/utils"
	"log"
	"os"
	"strings"
)

type ActionFunc func()

func ScanDirectory(path string) {
	log.Println("Scanning directory:", path)

	err := os.Chdir(path)
	if err != nil {
		log.Fatalf("Failed to change directory: %v", err)
	}

	configuration := ReadConfigFile()
	keyValues := ProcessConfigFile(configuration)
	RunScanners(keyValues)
}

func RunScanners(keyValues []string) {
	var lang string
	var actions = map[string]ActionFunc{
		"docker_compose":       func() { log.Println("action for docker_compose") },
		"dockerfile":           func() { log.Print(DockerfileExists(constants.CurrentPath)) },
		"readme":               func() { log.Print(Readme(constants.ReadmeFilePath)) },
		"ci":                   func() { log.Println("action for ci") },
		"cd":                   func() { log.Println("action for cd") },
		"conventional_commits": func() { log.Println("action for conventional_commits") },
		"pre_commit":           func() { log.Print(PreCommitExists(lang, constants.CurrentPath)) },
		"linter":               func() { log.Print(LinterJavascriptExists(constants.CurrentPath)) },
		"formatter":            func() { log.Println("action for formatter") },
		"unit":                 func() { log.Println("action for unit") },
		"integration":          func() { log.Println("action for integration") },
		"e2e":                  func() { log.Println("action for e2e") },
		"coverage":             func() { log.Println("action for coverage") },
		"stress":               func() { log.Println("action for stress") },
		"secrets":              func() { log.Println("action for secrets") },
		"iac":                  func() { log.Println("action for iac") },
		"code":                 func() { log.Println("action for code") },
		"container":            func() { log.Println("action for container") },
		"deps":                 func() { log.Println("action for deps") },
		"sast":                 func() { log.Println("action for sast") },
		"dast":                 func() { log.Println("action for dast") },
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

		if isActionEnabled(value) {
			action, exists := actions[key]
			if !exists {
				log.Printf("No action found for key '%s'\n", key)

				return
			}
			action()
		}
	}
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
