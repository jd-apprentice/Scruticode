package functions

import (
	"flag"
	"log"
)

// https://gobyexample.com/command-line-flags
func GenerateArguments(args []string) {
	fs := flag.NewFlagSet("generate", flag.ContinueOnError)
	lang := fs.String("languages", "golang", "Supported languages [golang, typescript, javascript, python].")
	platform := fs.String("platforms", "github", "Supported platforms [github, gitlab, azuredevops]")
	fs.Parse(args)

	log.Println("LANG:", *lang)
	log.Println("PLATFORM:", *platform)
}
