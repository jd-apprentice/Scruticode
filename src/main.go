package main

import (
	"Scruticode/src/config"
	"fmt"
)

func main() {
	configuration := config.ReadConfigFile()
	fmt.Print(configuration)
}
