package tests_test

import (
	"Scruticode/src/constants"
	"strings"
	"testing"
)

func isActionEnabled(action string) bool {
	return strings.ToLower(action) == constants.TrueAsString
}

func TestIsAction(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"enabled action", constants.TrueAsString, true},
		{"disabled action", constants.FalseAsString, false},
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
