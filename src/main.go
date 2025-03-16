package main

import (
	"Scruticode/src/config"
	"Scruticode/src/functions"
)

func main() {
	configuration := config.ReadConfigFile()
	functions.ProcessConfigFile(configuration)
}
