package utils

import "os"

func IfFileNotExists(file string, action func(string)) {
	if file == "" {
		return
	}
	if _, err := os.Stat(file); os.IsNotExist(err) {
		action(file)
	}
}
