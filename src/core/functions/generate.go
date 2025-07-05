package functions

import (
	"flag"
	"log"
)

// https://gobyexample.com/command-line-flags
func GenerateArguments(args []string) string {
	fs := flag.NewFlagSet("generate", flag.ContinueOnError)
	lang := fs.String("languages", "golang", "Supported languages [golang, typescript, javascript, python].")
	platform := fs.String("platforms", "github", "Supported platforms [github, gitlab, azuredevops]")
	folder := fs.String("folder", ".", "The folder to analyze.")
	fs.Parse(args)

	log.Println("LANG:", *lang)
	log.Println("PLATFORM:", *platform)
	log.Println("FOLDER:", *folder)

	return *folder
}
