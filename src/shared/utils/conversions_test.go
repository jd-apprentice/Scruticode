package utils

import (
	"testing"
)

func TestToAbsoluteString(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"string with only brackets", "[hello]", "hello"},
		{"string with only double quotes", `"hello"`, "hello"},
		{"string with no brackets or double quotes", "hello", "hello"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := ToAbsoluteString(tt.input)
			if actual != tt.expected {
				t.Errorf("ToAbsoluteString(%q) = %q, want %q", tt.input, actual, tt.expected)
			}
		})
	}
}
