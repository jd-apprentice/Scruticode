package functions

import (
	"log"
	"os"
)

// ExecuteScan orchestrates the scanning process based on configuration and arguments.
func ExecuteScan() {
	configuration := ReadConfigFile()
	ProcessConfigFile(configuration)
	_, _, directory, repository := GenerateArguments()

	if repository != "" {
		log.Println("Cloning repository:", repository)
		tempDir, err := CloneRepository(repository)
		if err != nil {
			log.Fatalf("Failed to clone repository: %v", err)
		}
		log.Println("Repository cloned to:", tempDir)
		defer RemoveTempDirectory(tempDir)
		ScanDirectory(tempDir)
	}

	if repository == "" {
		ScanDirectory(directory)
	}
}

func RemoveTempDirectory(tempDir string) {
	log.Println("Removing temporary directory:", tempDir)
	err := os.RemoveAll(tempDir)
	if err != nil {
		log.Printf("Failed to remove temporary directory: %v", err)
	}
}
