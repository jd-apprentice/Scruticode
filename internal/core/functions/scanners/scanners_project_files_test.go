package scanners

import (
	"os"
	"path/filepath"
	"testing"

	"Scruticode/internal/shared/constants"
)

func TestClineRulesExists(t *testing.T) {
	t.Parallel()

	t.Run("returns success when .clinerules directory exists and has files", func(t *testing.T) {
		t.Parallel()

		tempDir := t.TempDir()
		clineRulesDir := filepath.Join(tempDir, ".clinerules")
		if err := os.Mkdir(clineRulesDir, 0o755); err != nil {
			t.Fatalf("failed to create .clinerules directory: %v", err)
		}

		rulesFile := filepath.Join(clineRulesDir, "go.instructions.md")
		if err := os.WriteFile(rulesFile, []byte("# rules"), 0o600); err != nil {
			t.Fatalf("failed to create rule file: %v", err)
		}

		result := ClineRulesExists(clineRulesDir)
		if result.Status != constants.QualityCheckSuccess {
			t.Errorf("expected %v, got %v", constants.QualityCheckSuccess, result.Status)
		}
	})

	t.Run("returns failed when .clinerules directory does not exist", func(t *testing.T) {
		t.Parallel()

		result := ClineRulesExists("non-existent/.clinerules")
		if result.Status != constants.QualityCheckFailed {
			t.Errorf("expected %v, got %v", constants.QualityCheckFailed, result.Status)
		}
	})
}

func TestCopilotRulesExists(t *testing.T) {
	t.Parallel()

	t.Run("returns success when copilot instructions file exists", func(t *testing.T) {
		t.Parallel()
		tempDir := t.TempDir()
		path := filepath.Join(tempDir, "copilot-instructions.md")
		createFile(t, tempDir, "copilot-instructions.md", "# instructions")

		if CopilotRulesExists(path).Status != constants.QualityCheckSuccess {
			t.Fatal("expected CopilotRulesExists to return true")
		}
	})

	t.Run("returns failed when file does not exist", func(t *testing.T) {
		t.Parallel()
		if CopilotRulesExists("non-existent.md").Status != constants.QualityCheckFailed {
			t.Fatal("expected CopilotRulesExists to return false")
		}
	})
}

func TestDockerfileExists(t *testing.T) {
	t.Parallel()

	t.Run("returns success when Dockerfile exists", func(t *testing.T) {
		t.Parallel()
		tempDir := t.TempDir()
		createFile(t, tempDir, "Dockerfile", "FROM alpine")

		if DockerfileExists(tempDir).Status != constants.QualityCheckSuccess {
			t.Fatal("expected DockerfileExists to return true")
		}
	})

	t.Run("returns failed when Dockerfile does not exist", func(t *testing.T) {
		t.Parallel()
		tempDir := t.TempDir()

		if DockerfileExists(tempDir).Status != constants.QualityCheckFailed {
			t.Fatal("expected DockerfileExists to return false")
		}
	})
}

func TestReadme(t *testing.T) {
	t.Parallel()

	t.Run("returns success when file exists", func(t *testing.T) {
		t.Parallel()
		tempDir := t.TempDir()
		path := filepath.Join(tempDir, "README.md")
		createFile(t, tempDir, "README.md", "# project")

		if Readme(path).Status != constants.QualityCheckSuccess {
			t.Fatal("expected Readme to return true")
		}
	})

	t.Run("returns failed when file does not exist", func(t *testing.T) {
		t.Parallel()
		if Readme("non_existent_file.txt").Status != constants.QualityCheckFailed {
			t.Fatal("expected Readme to return false")
		}
	})
}
