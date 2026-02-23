package scanners

import (
	"Scruticode/internal/shared/constants"
	"testing"
)

func TestUnitTestsConfigured(t *testing.T) {
	t.Parallel()

	t.Run("returns true when unit folder exists", func(t *testing.T) {
		t.Parallel()
		tempDir := t.TempDir()
		createDir(t, tempDir, "unit")

		if UnitTestsConfigured(tempDir).Status != constants.QualityCheckSuccess {
			t.Fatal("expected UnitTestsConfigured to return true")
		}
	})

	t.Run("returns false without unit signals", func(t *testing.T) {
		t.Parallel()
		tempDir := t.TempDir()

		if UnitTestsConfigured(tempDir).Status == constants.QualityCheckSuccess {
			t.Fatal("expected UnitTestsConfigured to return false")
		}
	})
}

func TestIntegrationTestsConfigured(t *testing.T) {
	t.Parallel()
	tempDir := t.TempDir()
	createDir(t, tempDir, "tests/integration")

	if IntegrationTestsConfigured(tempDir).Status != constants.QualityCheckSuccess {
		t.Fatal("expected IntegrationTestsConfigured to return true")
	}
}

func TestE2ETestsConfigured(t *testing.T) {
	t.Parallel()
	tempDir := t.TempDir()
	createDir(t, tempDir, "cypress")

	if E2ETestsConfigured(tempDir).Status != constants.QualityCheckSuccess {
		t.Fatal("expected E2ETestsConfigured to return true")
	}
}

func TestCoverageConfigured(t *testing.T) {
	t.Parallel()
	tempDir := t.TempDir()
	createFile(t, tempDir, "codecov.yml", "coverage")

	if CoverageConfigured(tempDir).Status != constants.QualityCheckSuccess {
		t.Fatal("expected CoverageConfigured to return true")
	}
}

func TestStressTestsConfigured(t *testing.T) {
	t.Parallel()
	tempDir := t.TempDir()
	createDir(t, tempDir, "k6")

	if StressTestsConfigured(tempDir).Status != constants.QualityCheckSuccess {
		t.Fatal("expected StressTestsConfigured to return true")
	}
}
