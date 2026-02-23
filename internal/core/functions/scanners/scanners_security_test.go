package scanners

import (
	"Scruticode/internal/shared/constants"
	"testing"
)

func TestSecretScanningConfigured(t *testing.T) {
	t.Parallel()
	tempDir := t.TempDir()
	createFile(t, tempDir, ".gitleaks.toml", "title = 'gitleaks'")

	if SecretScanningConfigured(tempDir).Status != constants.QualityCheckSuccess {
		t.Fatal("expected SecretScanningConfigured to return true")
	}
}

func TestIACScanningConfigured(t *testing.T) {
	t.Parallel()
	tempDir := t.TempDir()
	createFile(t, tempDir, "main.tf", "terraform {}")

	if IACScanningConfigured(tempDir).Status != constants.QualityCheckSuccess {
		t.Fatal("expected IACScanningConfigured to return true")
	}
}

func TestCodeSecurityScanningConfigured(t *testing.T) {
	t.Parallel()
	tempDir := t.TempDir()
	createFile(t, tempDir, "sonar-project.properties", "sonar.projectKey=demo")

	if CodeSecurityScanningConfigured(tempDir).Status != constants.QualityCheckSuccess {
		t.Fatal("expected CodeSecurityScanningConfigured to return true")
	}
}

func TestDependencyScanningConfigured(t *testing.T) {
	t.Parallel()
	tempDir := t.TempDir()
	createFile(t, tempDir, ".github/dependabot.yml", "version: 2")

	if DependencyScanningConfigured(tempDir).Status != constants.QualityCheckSuccess {
		t.Fatal("expected DependencyScanningConfigured to return true")
	}
}

func TestSASTAndDASTConfigured(t *testing.T) {
	t.Parallel()
	t.Run("sast workflow", func(t *testing.T) {
		t.Parallel()
		tempDir := t.TempDir()
		createFile(t, tempDir, ".github/workflows/codeql-sast.yml", "name: codeql")

		if SASTConfigured(tempDir).Status != constants.QualityCheckSuccess {
			t.Fatal("expected SASTConfigured to return true")
		}
	})

	t.Run("dast workflow", func(t *testing.T) {
		t.Parallel()
		tempDir := t.TempDir()
		createFile(t, tempDir, ".github/workflows/zap-dast.yml", "name: zap")

		if DASTConfigured(tempDir).Status != constants.QualityCheckSuccess {
			t.Fatal("expected DASTConfigured to return true")
		}
	})
}
