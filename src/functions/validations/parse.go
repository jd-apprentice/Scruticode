package validations

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
			log.Println("action for golang")
		}
	case "python":
		{
			log.Println("action for python")
		}
	case "typescript":
	case "javascript":
	}
}

func platformHandler(platform string) {
	switch platform {
	case "github":
		{
			log.Println("action for github")
		}
	case "gitlab":
	case "azuredevops":
	}
}

func ExtraLangConfig(lang string) {
	var supportedLanguages = []string{"golang", "typescript", "javascript", "python"}
	extraConfig(lang, supportedLanguages, langHandler)
}

func ExtraPlatformConfig(platform string) {
	var supportedPlatforms = []string{"github", "gitlab", "azuredevops"}
	extraConfig(platform, supportedPlatforms, platformHandler)
}
