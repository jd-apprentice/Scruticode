package main

import (
	"Scruticode/src/config"
	"Scruticode/src/functions/arguments"
	"Scruticode/src/functions/core"
)

func main() {
	configuration := config.ReadConfigFile()
	core.ProcessConfigFile(configuration)
	arguments.Generate()
}
