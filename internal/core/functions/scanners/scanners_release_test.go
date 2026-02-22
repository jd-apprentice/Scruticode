package scanners

import (
	"Scruticode/internal/shared/constants"
	"testing"
)

func TestCIPipelineExists(t *testing.T) {
	t.Parallel()

	t.Run("returns true with gitlab ci file", func(t *testing.T) {
		t.Parallel()
		tempDir := t.TempDir()
		createFile(t, tempDir, ".gitlab-ci.yml", "stages: [test]")

		if CIPipelineExists(tempDir).Status != constants.QualityCheckSuccess {
			t.Fatal("expected CIPipelineExists to return true")
		}
	})

	t.Run("returns false when no ci signals exist", func(t *testing.T) {
		t.Parallel()
		tempDir := t.TempDir()

		if CIPipelineExists(tempDir).Status == constants.QualityCheckSuccess {
			t.Fatal("expected CIPipelineExists to return false")
		}
	})
}

func TestCDPipelineExists(t *testing.T) {
	t.Parallel()

	t.Run("returns true with deployment workflow", func(t *testing.T) {
		t.Parallel()
		tempDir := t.TempDir()
		createFile(t, tempDir, ".github/workflows/deploy.yml", "name: deploy")

		if CDPipelineExists(tempDir).Status != constants.QualityCheckSuccess {
			t.Fatal("expected CDPipelineExists to return true")
		}
	})

	t.Run("returns false with no cd configuration", func(t *testing.T) {
		t.Parallel()
		tempDir := t.TempDir()

		if CDPipelineExists(tempDir).Status == constants.QualityCheckSuccess {
			t.Fatal("expected CDPipelineExists to return false")
		}
	})
}
