package functions

import (
	"Scruticode/src/core/functions/options"
	"Scruticode/src/shared/constants"
	"Scruticode/src/shared/utils"
	"bytes"
	"log"
	"reflect"
	"strings"
	"testing"
)

type ActionFunc func()

func ProcessConfigFile(content string) {
	sections := extractSections(content)

	for _, section := range sections {
		_, keyValues := parseSection(section)
		// Check if is needed here to validate the HEADER, right now it's being ignored
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
	var actions = map[string]ActionFunc{
		"docker_compose":       func() { log.Println("action for docker_compose") },
		"dockerfile":           func() { log.Print(options.DockerfileExists()) },
		"readme":               func() { log.Print(options.Readme()) },
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

	const emptyAsString = ""
	for _, pair := range keyValues {
		key, value := parseKeyValuePair(pair)
		if key == emptyAsString {
			continue
		}

		keyAsString := utils.ToAbsoluteString(value)
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

func isSectionHeader(line string) bool {
	return strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]")
}

func isComment(line string) bool {
	return strings.HasPrefix(strings.TrimSpace(line), "#")
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

///// TEST SECTION /////

func TestIsAction(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"enabled action", "true", true},
		{"disabled action", "false", false},
		{"invalid action", "invalid", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := isActionEnabled(tt.input)
			if actual != tt.expected {
				t.Errorf("isActionEnabled(%q) = %v, want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestIsComment(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"empty string", "", false},
		{"string with whitespace", "   # comment", true},
		{"string starting with #", "# comment", true},
		{"string not starting with #", "comment", false},
		{"string with # in the middle", "comment # not a comment", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := isComment(tt.input)
			if actual != tt.expected {
				t.Errorf("isComment(%q) = %v, want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestIsSectionHeader(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"valid section header", "[section]", true},
		{"invalid section header (missing opening bracket)", "section]", false},
		{"invalid section header (missing closing bracket)", "[section", false},
		{"invalid section header (no brackets)", "section", false},
		{"invalid section header (only brackets)", "[]", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := isSectionHeader(tt.input)
			if actual != tt.expected {
				t.Errorf("isSectionHeader(%q) = %v, want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestParseKeyValuePair(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		input         string
		expectedKey   string
		expectedValue string
	}{
		{"valid key-value pair", "key=value", "key", "value"},
		{"invalid key-value pair", "key", "", ""},
		{"key-value pair with whitespaces", "  key  =  value  ", "key", "value"},
		// {"key-value pair with multiple equals signs", "key=value=value", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			key, value := parseKeyValuePair(tt.input)
			if key != tt.expectedKey || value != tt.expectedValue {
				t.Errorf("parseKeyValuePair(%q) = (%q, %q), want (%q, %q)", tt.input, key, value, tt.expectedKey, tt.expectedValue)
			}
		})
	}
}

func TestParseSection(t *testing.T) {
	t.Parallel()
	const headerString = "[header]"
	// structKeyValue := []string{"key1=value1, key2=value2"}
	tests := []struct {
		name      string
		section   string
		expected  string
		keyValues []string
	}{
		{"empty section", "", "", nil},
		{"section with only comments", "# comment\n# another comment", "", nil},
		{"section with header and no key-values", headerString, headerString, nil},
		{"section with header and key-values", "[header]\nkey=value", headerString, []string{"key=value"}},
		// {"section with multiple key-values", "[header]\nkey1=value1\nkey2=value2", headerString, structKeyValue},
		{"section with invalid key-values", "[header]\nkey", headerString, nil},
		// {"section with leading/trailing newlines", "\n[header]\nkey=value\n", headerString, []string{"key=value"}},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			header, keyValues := parseSection(testCase.section)
			if header != testCase.expected {
				t.Errorf("expected header %q, got %q", testCase.expected, header)
			}
			if !reflect.DeepEqual(keyValues, testCase.keyValues) {
				t.Errorf("expected key-values %v, got %v", testCase.keyValues, keyValues)
			}
		})
	}
}

func TestExtractSections(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		content  string
		expected []string
	}{
		{
			name:     "empty string",
			content:  "",
			expected: []string{""},
		},
		{
			name:     "single section",
			content:  "section1",
			expected: []string{"section1"},
		},
		{
			name:     "multiple sections",
			content:  "section1\n\nsection2\n\nsection3",
			expected: []string{"section1", "section2", "section3"},
		},
		{
			name:     "sections with trailing newlines",
			content:  "section1\n\nsection2\n\nsection3\n\n",
			expected: []string{"section1", "section2", "section3", ""},
		},
		{
			name:     "sections with leading newlines",
			content:  "\n\nsection1\n\nsection2\n\nsection3",
			expected: []string{"", "section1", "section2", "section3"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			actual := extractSections(test.content)
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("expected %v, got %v", test.expected, actual)
			}
		})
	}
}

func TestExtraLangConfig(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		lang     string
		expected bool
	}{
		{"unsupported language", "ruby", false},
		{"supported language", "golang", true},
		{"empty language", "", false},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			log.SetOutput(&bytes.Buffer{})
			var buf bytes.Buffer
			log.SetOutput(&buf)
			extraLangConfig(testCase.lang)
			if testCase.expected {
				if !strings.Contains(buf.String(), "action for "+testCase.lang) {
					t.Errorf("expected log message not found")
				}
			} else {
				if buf.String() != "" {
					t.Errorf("unexpected log message found")
				}
			}
		})
	}
}
