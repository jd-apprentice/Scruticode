package functions

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
)

func TestGenerateArguments(t *testing.T) {
	originalArgs := os.Args

	originalOutput := log.Writer()
	defer log.SetOutput(originalOutput)

	defer func() { os.Args = originalArgs }()

	t.Run("default arguments", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)

		os.Args = []string{"cmd"}

		GenerateArguments(os.Args[1:])

		output := buf.String()
		if !strings.Contains(output, "LANG: golang") {
			t.Errorf("Expected log to contain 'LANG: golang', but it didn't. Got: %s", output)
		}
		if !strings.Contains(output, "PLATFORM: github") {
			t.Errorf("Expected log to contain 'PLATFORM: github', but it didn't. Got: %s", output)
		}
	})

	t.Run("with provided arguments", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)

		os.Args = []string{"cmd", "-languages=typescript", "-platforms=gitlab"}

		GenerateArguments(os.Args[1:])

		output := buf.String()
		if !strings.Contains(output, "LANG: typescript") {
			t.Errorf("Expected log to contain 'LANG: typescript', but it didn't. Got: %s", output)
		}
		if !strings.Contains(output, "PLATFORM: gitlab") {
			t.Errorf("Expected log to contain 'PLATFORM: gitlab', but it didn't. Got: %s", output)
		}
	})
}
