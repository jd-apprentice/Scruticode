package tests_test

import (
	"strings"
	"testing"
)

func isSectionHeader(line string) bool {
	return strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]")
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
