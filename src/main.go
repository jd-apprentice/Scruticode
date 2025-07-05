package main

import (
	"Scruticode/src/core/functions"
	"os"
)

func main() {
	functions.Init()
	folder := functions.GenerateArguments(os.Args[1:])
	configuration := functions.ReadConfigFile()
	functions.ProcessConfigFile(configuration, folder)
}
