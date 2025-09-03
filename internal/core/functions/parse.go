package functions

import (
	"log"
	"slices"
)

func isValueAllowed(value string, allowedValues []string) bool {
	return slices.Contains(allowedValues, value)
}

func extraConfig(value string, allowedValues []string, handler func(string)) {
	if !isValueAllowed(value, allowedValues) {
		return
	}
	handler(value)
}

func langHandler(lang string) {
	switch lang {
	case "golang":
		{
			log.Println("Using configuration for golang")
		}
	case "python":
		{
			log.Println("action for python")
		}
	case "typescript":
		{
			log.Println("action for typescript")
		}
	case "javascript":
		{
			log.Println("action for javascript")
		}
	}
}

func platformHandler(platform string) {
	switch platform {
	case "github":
		{
			log.Println("Using configuration for github")
		}
	case "gitlab":
		{
			log.Println("action for gitlab")
		}
	case "azuredevops":
		{
			log.Println("action for azuredevops")
		}
	}
}

func extraLangConfig(lang string) {
	var supportedLanguages = []string{"golang", "typescript", "javascript", "python"}
	extraConfig(lang, supportedLanguages, langHandler)
}

func extraPlatformConfig(platform string) {
	var supportedPlatforms = []string{"github", "gitlab", "azuredevops"}
	extraConfig(platform, supportedPlatforms, platformHandler)
}
