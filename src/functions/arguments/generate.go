package arguments

import (
	"flag"
	"log"
)

// https://gobyexample.com/command-line-flags
func Generate() {

	lang := flag.String("languages", "golang", "Supported languages [golang, typescript, javascript, python].")
	platform := flag.String("platforms", "github", "Supported platforms [github, gitlab, azuredevops]")
	flag.Parse()

	log.Println("LANG:", *lang)
	log.Println("PLATFORM:", *platform)
}
