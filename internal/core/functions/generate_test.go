package functions

import (
	"bytes"
	"flag"
	"log"
	"os"
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

		fs := flag.NewFlagSet("test", flag.ContinueOnError)
		fs.SetOutput(&buf)

		os.Args = []string{"cmd"}

		lang, platform, _, _ := GenerateArguments(fs)

		if lang != "golang" {
			t.Errorf("Expected language to be 'golang', but got: %s", lang)
		}
		if platform != "github" {
			t.Errorf("Expected platform to be 'github', but got: %s", platform)
		}
	})

	t.Run("with provided arguments", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)

		fs := flag.NewFlagSet("test", flag.ContinueOnError)
		fs.SetOutput(&buf) // Redirect flag output to buffer

		os.Args = []string{"cmd", "-languages=typescript", "-platforms=gitlab"} // Simulate command line arguments

		lang, platform, _, _ := GenerateArguments(fs)

		if lang != "typescript" {
			t.Errorf("Expected language to be 'typescript', but got: %s", lang)
		}
		if platform != "gitlab" {
			t.Errorf("Expected platform to be 'gitlab', but got: %s", platform)
		}
	})
}
