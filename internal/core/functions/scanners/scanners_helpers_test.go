package scanners

import "testing"

func TestHelpersBasicBehavior(t *testing.T) {
	t.Parallel()
	tempDir := t.TempDir()

	createDir(t, tempDir, "nested")
	createFile(t, tempDir, "nested/file_test.go", "package sample")

	if !directoryExists(tempDir, "nested") {
		t.Fatal("expected directoryExists to return true")
	}

	if !anyFileExists(tempDir, "nested/file_test.go") {
		t.Fatal("expected anyFileExists to return true")
	}

	if !hasFileWithSuffix(tempDir, "_test.go") {
		t.Fatal("expected hasFileWithSuffix to return true")
	}
}
