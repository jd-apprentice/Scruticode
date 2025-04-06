package main

import "Scruticode/src/core/functions"

func main() {
	configuration := functions.ReadConfigFile()
	functions.ProcessConfigFile(configuration)
	functions.GenerateArguments()
}
