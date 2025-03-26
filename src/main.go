package main

import (
	"Scruticode/src/config"
	"Scruticode/src/functions"
	"Scruticode/src/functions/arguments"
)

// https://stackoverflow.com/questions/56039154/is-it-really-bad-to-use-init-functions-in-go#56039373
// https://leighmcculloch.com/posts/tool-go-check-no-globals-no-inits/
func main() {
	configuration := config.ReadConfigFile()
	functions.ProcessConfigFile(configuration)
	arguments.Generate()
}
