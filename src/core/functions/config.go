package functions

import (
	"Scruticode/src/shared/constants"
	"log"
	"os"
	"os/user"
)

// This function should read a file located in $HOME/.config/scruticode/settings.toml
// If the file is not existing, it should create it
func ReadConfigFile() string {
	usr, errUsr := user.Current()
	if errUsr != nil {
		log.Println(constants.ErrMessageUser, errUsr)

		return ""
	}

	filePath := usr.HomeDir + constants.ConfigFilePath
	if _, errPath := os.Stat(filePath); os.IsNotExist(errPath) {
		InitConfigFile(usr.HomeDir)
	}

	content, errReadFile := os.ReadFile(filePath)
	if errReadFile != nil {
		log.Println(constants.ErrMessageReading, errReadFile)
	}

	return string(content)
}
