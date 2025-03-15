package config

import (
	"Scruticode/src/constants"
	"log"
	"os"
	"os/user"
	"path/filepath"
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
	dirPath := filepath.Dir(filePath)

	if _, errPath := os.Stat(filePath); os.IsNotExist(errPath) {
		errCreateFolder := os.MkdirAll(dirPath, constants.DefaultFilePermissions)
		if errCreateFolder != nil {
			log.Println(constants.ErrMessageDirectory, errCreateFolder)

			return ""
		}

		file, errCreateFile := os.Create(filePath)
		if errCreateFile != nil {
			log.Println(constants.ErrMessageFile, errCreateFile)

			return ""
		}
		defer file.Close()
	}

	content, errReadFile := os.ReadFile(filePath)
	if errReadFile != nil {
		log.Println(constants.ErrMessageReading, errReadFile)
	}

	return string(content)
}
