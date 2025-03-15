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
	usr, err := user.Current()
	if err != nil {
		log.Println(constants.ErrMessageUser, err)

		return ""
	}

	filePath := usr.HomeDir + constants.ConfigFilePath
	dirPath := filepath.Dir(filePath)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, constants.DefaultFilePermissions)
		if err != nil {
			log.Println(constants.ErrMessageDirectory, err)

			return ""
		}

		file, err := os.Create(filePath)
		if err != nil {
			log.Println(constants.ErrMessageFile, err)

			return ""
		}
		defer file.Close()
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Println(constants.ErrMessageReading, err)
	}

	return string(content)
}
