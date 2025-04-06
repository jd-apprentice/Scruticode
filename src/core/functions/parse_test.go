package functions

import (
	"bytes"
	"log"
	"strings"
	"testing"
)

func TestExtraLangConfig(t *testing.T) {
	tests := []struct {
		name     string
		lang     string
		expected bool
	}{
		{"supported language", "golang", true},
		{"unsupported language", "ruby", false},
		{"empty language", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			log.SetOutput(&buf)
			extraLangConfig(tt.lang)
			if tt.expected {
				if !strings.Contains(buf.String(), "action for "+tt.lang) {
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

func TestExtraPlatformConfig(t *testing.T) {
	tests := []struct {
		name     string
		platform string
		expected string
	}{
		{"supported platform", "github", "action for github"},
		{"unsupported platform", "bitbucket", ""},
		{"empty platform string", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := &bytes.Buffer{}
			log.SetOutput(buf)
			extraPlatformConfig(tt.platform)
			actual := buf.String()
			if tt.expected != "" {
				if !strings.Contains(actual, tt.expected) {
					t.Errorf("extraPlatformConfig(%q) = %q, want to contain %q", tt.platform, actual, tt.expected)
				}
			} else {
				if actual != "" {
					t.Errorf("extraPlatformConfig(%q) = %q, want %q", tt.platform, actual, tt.expected)
				}
			}
		})
	}
}
