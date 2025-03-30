package validations_test

import (
	"Scruticode/src/functions/validations"
	"bytes"
	"log"
	"strings"
	"testing"
)

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
			validations.ExtraLangConfig(testCase.lang)
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
