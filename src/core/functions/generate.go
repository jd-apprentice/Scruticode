package functions

import (
	"flag"
	"log"
)

// https://gobyexample.com/command-line-flags
func GenerateArguments() (string, string, string, string) {
	lang := flag.String("languages", "golang", "Supported languages [golang, typescript, javascript, python].")
	platform := flag.String("platforms", "github", "Supported platforms [github, gitlab, azuredevops]")
	directory := flag.String("directory", ".", "The directory to scan.")
	repository := flag.String("repository", "", "The repository to scan.")
	flag.Parse()

	log.Println("LANG:", *lang)
	log.Println("PLATFORM:", *platform)
	log.Println("DIRECTORY:", *directory)
	log.Println("REPOSITORY:", *repository)

	return *lang, *platform, *directory, *repository
}
