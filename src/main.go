package main

import (
	"Scruticode/src/core/functions"
	"os"
)

func main() {
	functions.Init()
	configuration := functions.ReadConfigFile()
	functions.ProcessConfigFile(configuration)
	functions.GenerateArguments(os.Args[1:])
}
