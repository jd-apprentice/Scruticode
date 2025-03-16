package functions

import (
	"Scruticode/src/constants"
	"Scruticode/src/types"
	"log"
	"strings"
)

var actions = map[string]types.ActionFunc{
	"docker_compose":       func() { log.Println("action for docker_compose") },
	"dockerfile":           func() { log.Println("action for dockerfile") },
	"readme":               func() { log.Println("action for readme") },
	"ci":                   func() { log.Println("action for ci") },
	"cd":                   func() { log.Println("action for cd") },
	"conventional_commits": func() { log.Println("action for conventional_commits") },
	"pre_commit":           func() { log.Println("action for pre_commit") },
	"linter":               func() { log.Println("action for linter") },
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

func ProcessConfigFile(content string) {
	sections := extractSections(content)

	for _, section := range sections {
		_, keyValues := parseSection(section)
		// TODO: Check if is needed here to validate the HEADER, right now it's being ignored
		processKeyValues(keyValues)
	}
}

func extractSections(content string) []string {
	return strings.Split(content, "\n\n")
}

func parseSection(section string) (string, []string) {
	lines := strings.Split(section, "\n")
	if len(lines) == 0 {
		return "", nil
	}

	var keyValues []string
	var header string

	for index, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == "" || isComment(trimmedLine) {
			continue
		}

		if index == constants.IsEmpty && isSectionHeader(trimmedLine) {
			header = trimmedLine

			continue
		}

		if strings.Contains(trimmedLine, "=") {
			keyValues = append(keyValues, trimmedLine)
		}
	}

	return header, keyValues
}

func processKeyValues(keyValues []string) {
	for _, pair := range keyValues {
		key, value := parseKeyValuePair(pair)
		if key == "" {
			continue
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

func isSectionHeader(line string) bool {
	return strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]")
}

func isComment(line string) bool {
	return strings.HasPrefix(strings.TrimSpace(line), "#")
}

func isActionEnabled(action string) bool {
	return strings.ToLower(action) == constants.TrueAsString
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
