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
			log.Println("TODO: Action for Golang")
		}
	case "python":
		{
			log.Println("TODO: Action for Python")
		}
	case "typescript":
		{
			log.Println("TODO: Action for Typescript")
		}
	case "javascript":
		{
			log.Println("TODO: Action for Javascript")
		}
	}
}

func platformHandler(platform string) {
	switch platform {
	case "github":
		{
			log.Println("TODO: Action for GitHub")
		}
	case "gitlab":
		{
			log.Println("TODO: Action for GitLab")
		}
	case "azuredevops":
		{
			log.Println("TODO: Action for Azure DevOps")
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
