package utils

import "log"

func LoggerError(message, file string) {
	log.Fatalf("%s: %s\n", message, file)
}

func LoggerDebug(message string, file string) {
	log.Printf("%s: %s\n", message, file)
}
