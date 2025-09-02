package functions

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
)

func CloneRepository(repoURL string) (string, error) {
	tempDir, err := os.MkdirTemp("", "scruticode-")
	if err != nil {
		return "", fmt.Errorf("failed to create temporary directory: %w", err)
	}

	_, err = git.PlainClone(tempDir, false, &git.CloneOptions{
		URL:      repoURL,
		Progress: nil,
	})
	if err != nil {
		return "", fmt.Errorf("failed to clone repository: %w", err)
	}

	return tempDir, nil
}
