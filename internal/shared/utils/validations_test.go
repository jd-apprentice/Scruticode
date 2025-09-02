package utils

import (
	"log"
	"os"
	"testing"
)

func setupSuite(tb testing.TB) func(tb testing.TB) {
	tb.Helper()
	log.Println("Setup suite", tb.Name())

	return func(_ testing.TB) {
		os.Remove("existing_file.txt")
	}
}

func TestIfFileNotExists(t *testing.T) {
	t.Parallel()
	teardown := setupSuite(t)
	defer teardown(t)

	type test struct {
		name     string
		filePath string
		expected bool
	}

	tests := []test{
		{
			name:     "file does not exist",
			filePath: "non_existent_file.txt",
			expected: true,
		},
		{
			name:     "file exists",
			filePath: "existing_file.txt",
			expected: false,
		},
		{
			name:     "empty file path",
			filePath: "",
			expected: false,
		},
		{
			name:     "file path is a directory",
			filePath: "./",
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			if test.name == "file exists" {
				os.WriteFile(test.filePath, []byte("hello"), 0644)
				defer os.Remove(test.filePath)
			}

			var executed bool
			IfFileNotExists(test.filePath, func(string) {
				executed = true
			})
			if executed != test.expected {
				t.Errorf("expected %t, got %t", test.expected, executed)
			}
		})
	}
}
