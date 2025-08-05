package functions

import (
	"Scruticode/src/shared/constants"
	"strings"
)

func ProcessConfigFile(content string) []string {
	sections := extractSections(content)

	var allKeyValues []string
	for _, section := range sections {
		_, keyValues := parseSection(section)
		allKeyValues = append(allKeyValues, keyValues...)
	}
	return allKeyValues
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

func isSectionHeader(line string) bool {
	return strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]")
}

func isComment(line string) bool {
	return strings.HasPrefix(strings.TrimSpace(line), "#")
}
