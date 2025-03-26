package utils

import "strings"

func ToAbsoluteString(value string) string {
	removeBrackets := strings.Trim(value, "[]")
	removeDobleQuotes := strings.Trim(removeBrackets, `"`)
	stringValue := strings.TrimSpace(removeDobleQuotes)

	return stringValue
}
