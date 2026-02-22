package scanners

import (
	"Scruticode/internal/shared/constants"
	"testing"
)

func TestConventionalCommitsConfigured(t *testing.T) {
	t.Parallel()

	t.Run("returns true with commitlint file", func(t *testing.T) {
		t.Parallel()
		tempDir := t.TempDir()
		createFile(t, tempDir, "commitlint.config.js", "module.exports = {}")

		if ConventionalCommitsConfigured(tempDir).Status != constants.QualityCheckSuccess {
			t.Fatal("expected ConventionalCommitsConfigured to return true")
		}
	})

	t.Run("returns false when config is absent", func(t *testing.T) {
		t.Parallel()
		tempDir := t.TempDir()

		if ConventionalCommitsConfigured(tempDir).Status == constants.QualityCheckSuccess {
			t.Fatal("expected ConventionalCommitsConfigured to return false")
		}
	})
}

func TestFormatterConfigured(t *testing.T) {
	t.Parallel()

	t.Run("returns true with prettier config", func(t *testing.T) {
		t.Parallel()
		tempDir := t.TempDir()
		createFile(t, tempDir, ".prettierrc", "{}")

		if FormatterConfigured(tempDir).Status != constants.QualityCheckSuccess {
			t.Fatal("expected FormatterConfigured to return true")
		}
	})

	t.Run("returns false with no formatter signal", func(t *testing.T) {
		t.Parallel()
		tempDir := t.TempDir()

		if FormatterConfigured(tempDir).Status == constants.QualityCheckSuccess {
			t.Fatal("expected FormatterConfigured to return false")
		}
	})
}
