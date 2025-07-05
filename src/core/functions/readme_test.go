package functions

import (
	"os"
	"testing"

	"Scruticode/src/shared/constants"
)

func TestReadme(t *testing.T) {
	t.Run("Readme file exists", func(t *testing.T) {
		t.Parallel()
		tmpFile, err := os.CreateTemp("", "README.md")
		if err != nil {
			t.Fatal(err)
		}
		defer os.Remove(tmpFile.Name())

		resp := Readme(tmpFile.Name())
		if resp.Status != constants.QualityCheckSuccess {
			t.Errorf("Expected success, got %v", resp.Status)
		}
	})

	t.Run("Readme file does not exist", func(t *testing.T) {
		t.Parallel()
		resp := Readme("non_existent_file.txt")
		if resp.Status != constants.QualityCheckFailed {
			t.Errorf("Expected failure, got %v", resp.Status)
		}
	})
}
