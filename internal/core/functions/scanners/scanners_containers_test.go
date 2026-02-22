package scanners

import (
	"Scruticode/internal/shared/constants"
	"testing"
)

func TestDockerComposeExists(t *testing.T) {
	t.Parallel()

	t.Run("returns true when docker compose file exists", func(t *testing.T) {
		t.Parallel()

		tempDir := t.TempDir()
		createFile(t, tempDir, "docker-compose.yml", "version: '3'")

		if DockerComposeExists(tempDir).Status != constants.QualityCheckSuccess {
			t.Fatal("expected DockerComposeExists to return true")
		}
	})

	t.Run("returns false when no compose file exists", func(t *testing.T) {
		t.Parallel()

		tempDir := t.TempDir()

		if DockerComposeExists(tempDir).Status == constants.QualityCheckSuccess {
			t.Fatal("expected DockerComposeExists to return false")
		}
	})
}

func TestContainerSecurityScanningConfigured(t *testing.T) {
	t.Parallel()

	t.Run("returns true when docker context and workflow are present", func(t *testing.T) {
		t.Parallel()

		tempDir := t.TempDir()
		createFile(t, tempDir, "Dockerfile", "FROM alpine")
		createFile(t, tempDir, ".github/workflows/trivy-container.yml", "name: trivy")

		if ContainerSecurityScanningConfigured(tempDir).Status != constants.QualityCheckSuccess {
			t.Fatal("expected ContainerSecurityScanningConfigured to return true")
		}
	})

	t.Run("returns false when docker context is missing", func(t *testing.T) {
		t.Parallel()

		tempDir := t.TempDir()
		createFile(t, tempDir, ".github/workflows/trivy-container.yml", "name: trivy")

		if ContainerSecurityScanningConfigured(tempDir).Status == constants.QualityCheckSuccess {
			t.Fatal("expected ContainerSecurityScanningConfigured to return false")
		}
	})
}
