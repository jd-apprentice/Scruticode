package functions

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestInitConfigFile(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "test-init")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	originalHome, isset := os.LookupEnv("HOME")
	os.Setenv("HOME", tmpDir)
	if !isset {
		defer os.Unsetenv("HOME")
	}
	defer os.Setenv("HOME", originalHome)

	t.Run("config file exists", func(t *testing.T) {
		configDir := filepath.Join(tmpDir, ".config", "scruticode")
		if err := os.MkdirAll(configDir, 0755); err != nil {
			t.Fatalf("Failed to create config dir: %v", err)
		}
		configFilePath := filepath.Join(configDir, "settings.toml")
		if _, err := os.Create(configFilePath); err != nil {
			t.Fatalf("Failed to create config file: %v", err)
		}

		InitConfigFile(tmpDir)
	})

	t.Run("config file does not exist", func(t *testing.T) {
		configFilePath := filepath.Join(tmpDir, ".config", "scruticode", "settings.toml")
		os.Remove(configFilePath)

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("test config"))
		}))
		defer server.Close()

		originalURL := exampleConfigURL
		exampleConfigURL = server.URL
		defer func() { exampleConfigURL = originalURL }()

		InitConfigFile(tmpDir)

		if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
			t.Error("Expected config file to be created, but it was not")
		}
	})
}

func TestCreateConfigFile(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "test-create-config")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	t.Run("successful creation", func(t *testing.T) {
		configFilePath := filepath.Join(tmpDir, "settings.toml")

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("test config"))
		}))
		defer server.Close()

		originalURL := exampleConfigURL
		exampleConfigURL = server.URL
		defer func() { exampleConfigURL = originalURL }()

		createConfigFile(configFilePath)

		content, err := os.ReadFile(configFilePath)
		if err != nil {
			t.Fatalf("Failed to read created config file: %v", err)
		}
		if string(content) != "test config" {
			t.Errorf("Expected config file content to be 'test config', got '%s'", string(content))
		}
	})

	t.Run("failed to create file", func(t *testing.T) {
		configFilePath := "/dev/null/settings.toml"
		createConfigFile(configFilePath)
	})

	t.Run("http request fails", func(t *testing.T) {
		configFilePath := filepath.Join(tmpDir, "settings.toml")

		originalURL := exampleConfigURL
		exampleConfigURL = "http://invalid-url"
		defer func() { exampleConfigURL = originalURL }()

		createConfigFile(configFilePath)
	})
}
