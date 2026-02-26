package functions

import (
	"flag"
	"fmt"
	"os"

	"Scruticode/internal/shared/version"
)

// https://gobyexample.com/command-line-flags
func GenerateArguments(fs *flag.FlagSet) (string, string, string, string) {
	for _, arg := range os.Args[1:] {
		if arg == "-v" || arg == "--version" {
			version := version.GetVersion()
			fmt.Printf("Scruticode version %s\n", version)
			os.Exit(0)
		}
	}

	lang := fs.String("languages", "golang", "Supported languages [golang, typescript, javascript, python].")
	platform := fs.String("platforms", "github", "Supported platforms [github, gitlab, azuredevops]")
	directory := fs.String("directory", ".", "The directory to scan.")
	repository := fs.String("repository", "", "The repository to scan.")
	fs.Parse(os.Args[1:])

	return *lang, *platform, *directory, *repository
}
