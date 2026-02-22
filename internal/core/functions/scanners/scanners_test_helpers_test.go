package scanners

import (
	"os"
	"path/filepath"
	"testing"
)

func createFile(t *testing.T, basePath, relativePath, content string) {
	t.Helper()

	absolutePath := filepath.Join(basePath, relativePath)
	if err := os.MkdirAll(filepath.Dir(absolutePath), 0o755); err != nil {
		t.Fatalf("failed to create dir for %s: %v", relativePath, err)
	}

	if err := os.WriteFile(absolutePath, []byte(content), 0o600); err != nil {
		t.Fatalf("failed to create file %s: %v", relativePath, err)
	}
}

func createDir(t *testing.T, basePath, relativePath string) {
	t.Helper()

	if err := os.MkdirAll(filepath.Join(basePath, relativePath), 0o755); err != nil {
		t.Fatalf("failed to create directory %s: %v", relativePath, err)
	}
}
