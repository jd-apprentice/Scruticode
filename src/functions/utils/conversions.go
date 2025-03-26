package utils

import "strings"

func ToAbsoluteString(value string) string {
	removeBrackets := strings.Trim(value, "[]")
	removeDobleQuotes := strings.Trim(removeBrackets, `"`)
	stringLang := strings.TrimSpace(removeDobleQuotes)

	return stringLang
}
