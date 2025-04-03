package utils

import "log"

func LoggerErrorFile(message, file string) {
	log.Fatalf("%s: %s\n", message, file)
}

func LoggerDebugFile(message string, file string) {
	log.Printf("%s: %s\n", message, file)
}
