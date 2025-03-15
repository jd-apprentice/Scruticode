package main

import (
	"Scruticode/src/config"
	"log"
)

func main() {
	configuration := config.ReadConfigFile()
	log.Print(configuration)
}
