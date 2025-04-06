package functions

import (
	"reflect"
	"testing"
)

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
		{"section with invalid key-values", "[header]\nkey", headerString, nil},
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
