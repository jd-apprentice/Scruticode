package tests_test

import (
	"strings"
	"testing"
)

func isComment(line string) bool {
	return strings.HasPrefix(strings.TrimSpace(line), "#")
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
