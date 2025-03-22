package functions

import (
	"Scruticode/src/constants"
	"Scruticode/src/functions/utils"
	"context"
	"io"
	"log"
	"net/http"
	"os"
)

func InitConfigFile() {
	homeDir, errHomeUsr := os.UserHomeDir()
	if errHomeUsr != nil {
		log.Fatal(errHomeUsr)
	}

	configFile := homeDir + constants.ConfigFilePath
	utils.IfFileNotExists(configFile, createConfigFile)

	info, errFileNotExists := os.Stat(configFile)
	if errFileNotExists != nil {
		log.Println(errFileNotExists)
	}

	if info.Size() == constants.IsEmpty {
		os.Remove(configFile)
		createConfigFile(configFile)
	}
}

func createConfigFile(configFilePath string) {
	file, errFailedToCreate := os.Create(configFilePath)
	if errFailedToCreate != nil {
		log.Println(constants.ErrMessageFile)

		return
	}
	defer file.Close()

	ctx := context.Background()
	request, errHTTP := http.NewRequestWithContext(ctx, http.MethodGet, constants.ExampleConfig, nil)
	if errHTTP != nil {
		log.Println(errHTTP)

		return
	}

	response, errDo := http.DefaultClient.Do(request)
	if errDo != nil {
		log.Println(errDo)

		return
	}
	defer response.Body.Close()

	if _, errCopyContent := io.Copy(file, response.Body); errCopyContent != nil {
		log.Println(errCopyContent)

		return
	}

	usrHomeDir, _ := os.UserHomeDir()
	log.Println(usrHomeDir + constants.ConfigFilePath + " " + constants.FileCreated)
}
