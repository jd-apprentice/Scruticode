package main

import (
	"Scruticode/src/config"
	"Scruticode/src/functions"
	"Scruticode/src/functions/arguments"
)

func main() {
	configuration := config.ReadConfigFile()
	functions.ProcessConfigFile(configuration)
	arguments.Generate()
}
