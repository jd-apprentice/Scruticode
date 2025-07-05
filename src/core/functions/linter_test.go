package functions

import (
	"Scruticode/src/shared/constants"
	"os"
	"path/filepath"
	"testing"
)

func TestLinterJavascriptExists(t *testing.T) {

	const packageFile = "package.json"

	tmpDir, err := os.MkdirTemp("", "test-linter")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	t.Run("valid package.json", func(t *testing.T) {
		packageJSONContent := `{
			"dependencies": {
				"eslint": "8.0.0"
			},
			"scripts": {
				"lint": "eslint ."
			}
		}`
		packageJSONPath := filepath.Join(tmpDir, "package.json")
		if err := os.WriteFile(packageJSONPath, []byte(packageJSONContent), 0644); err != nil {
			t.Fatalf("Failed to write package.json: %v", err)
		}

		response := LinterJavascriptExists(tmpDir)
		if response.Status != constants.QualityCheckSuccess {
			t.Errorf("Expected status %d, got %d", constants.QualityCheckSuccess, response.Status)
		}
	})

	t.Run("missing eslint dependency", func(t *testing.T) {
		packageJSONContent := `{
			"dependencies": {},
			"scripts": {
				"lint": "eslint ."
			}
		}`
		packageJSONPath := filepath.Join(tmpDir, "package.json")
		if err := os.WriteFile(packageJSONPath, []byte(packageJSONContent), 0644); err != nil {
			t.Fatalf("Failed to write package.json: %v", err)
		}

		response := LinterJavascriptExists(tmpDir)
		if response.Status != constants.QualityCheckFailed {
			t.Errorf("Expected status %d, got %d", constants.QualityCheckFailed, response.Status)
		}
	})

	t.Run("missing lint script", func(t *testing.T) {
		packageJSONContent := `{
			"dependencies": {
				"eslint": "8.0.0"
			},
			"scripts": {}
		}`
		packageJSONPath := filepath.Join(tmpDir, "package.json")
		if err := os.WriteFile(packageJSONPath, []byte(packageJSONContent), 0644); err != nil {
			t.Fatalf("Failed to write package.json: %v", err)
		}

		response := LinterJavascriptExists(tmpDir)
		if response.Status != constants.QualityCheckFailed {
			t.Errorf("Expected status %d, got %d", constants.QualityCheckFailed, response.Status)
		}
	})

	t.Run("invalid json", func(t *testing.T) {
		packageJSONContent := `{
			"dependencies": {
				"eslint": "8.0.0"
			},
			"scripts": {
				"lint": "eslint ."
			},
		}`
		packageJSONPath := filepath.Join(tmpDir, "package.json")
		if err := os.WriteFile(packageJSONPath, []byte(packageJSONContent), 0644); err != nil {
			t.Fatalf("Failed to write package.json: %v", err)
		}

		response := LinterJavascriptExists(tmpDir)
		if response.Status != constants.QualityCheckFailed {
			t.Errorf("Expected status %d, got %d", constants.QualityCheckFailed, response.Status)
		}
	})

	t.Run("missing package.json", func(t *testing.T) {
		os.Remove(filepath.Join(tmpDir, "package.json"))
		response := LinterJavascriptExists(tmpDir)
		if response.Status != constants.QualityCheckFailed {
			t.Errorf("Expected status %d, got %d", constants.QualityCheckFailed, response.Status)
		}
	})
}

func TestHasDependency(t *testing.T) {
	tests := []struct {
		name        string
		packageJSON map[string]interface{}
		dependency  string
		expected    bool
	}{
		{
			name: "dependency exists",
			packageJSON: map[string]interface{}{
				"dependencies": map[string]interface{}{
					"react": "17.0.2",
				},
			},
			dependency: "react",
			expected:   true,
		},
		{
			name: "dependency does not exist",
			packageJSON: map[string]interface{}{
				"dependencies": map[string]interface{}{
					"vue": "3.2.37",
				},
			},
			dependency: "react",
			expected:   false,
		},
		{
			name: "dev dependency exists",
			packageJSON: map[string]interface{}{
				"devDependencies": map[string]interface{}{
					"eslint": "8.0.0",
				},
			},
			dependency: "eslint",
			expected:   true,
		},
		{
			name:        "no dependencies",
			packageJSON: map[string]interface{}{},
			dependency:  "react",
			expected:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasDependency(tt.packageJSON, tt.dependency); got != tt.expected {
				t.Errorf("hasDependency() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestHasScript(t *testing.T) {
	tests := []struct {
		name        string
		packageJSON map[string]interface{}
		script      string
		expected    bool
	}{
		{
			name: "script exists",
			packageJSON: map[string]interface{}{
				"scripts": map[string]interface{}{
					"start": "react-scripts start",
				},
			},
			script:   "start",
			expected: true,
		},
		{
			name: "script does not exist",
			packageJSON: map[string]interface{}{
				"scripts": map[string]interface{}{
					"build": "react-scripts build",
				},
			},
			script:   "start",
			expected: false,
		},
		{
			name:        "no scripts",
			packageJSON: map[string]interface{}{},
			script:      "start",
			expected:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasScript(tt.packageJSON, tt.script); got != tt.expected {
				t.Errorf("hasScript() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestReadAndParsePackageJSON(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "test-read-parse")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	t.Run("valid package.json", func(t *testing.T) {
		content := `{"name": "test-project", "version": "1.0.0"}`
		path := filepath.Join(tmpDir, "package.json")
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			t.Fatalf("Failed to write package.json: %v", err)
		}

		data, err := readAndParsePackageJSON(path)
		if err != nil {
			t.Errorf("readAndParsePackageJSON() error = %v, wantErr %v", err, false)
		}
		if data["name"] != "test-project" {
			t.Errorf("Expected name to be 'test-project', got '%s'", data["name"])
		}
	})

	t.Run("invalid package.json", func(t *testing.T) {
		content := `{"name": "test-project", "version": "1.0.0",}`
		path := filepath.Join(tmpDir, "package.json")
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			t.Fatalf("Failed to write package.json: %v", err)
		}

		_, err := readAndParsePackageJSON(path)
		if err == nil {
			t.Errorf("readAndParsePackageJSON() error = %v, wantErr %v", err, true)
		}
	})

	t.Run("file not found", func(t *testing.T) {
		_, err := readAndParsePackageJSON(filepath.Join(tmpDir, "nonexistent.json"))
		if err == nil {
			t.Errorf("readAndParsePackageJSON() error = %v, wantErr %v", err, true)
		}
	})
}

func TestGetBasePath(t *testing.T) {
	t.Run("with project path", func(t *testing.T) {
		path, err := getBasePath("/tmp")
		if err != nil {
			t.Errorf("getBasePath() error = %v, wantErr %v", err, false)
		}
		if path != "/tmp" {
			t.Errorf("Expected path to be '/tmp', got '%s'", path)
		}
	})

	t.Run("without project path", func(t *testing.T) {
		wd, err := os.Getwd()
		if err != nil {
			t.Fatalf("Failed to get current working directory: %v", err)
		}
		path, err := getBasePath()
		if err != nil {
			t.Errorf("getBasePath() error = %v, wantErr %v", err, false)
		}
		if path != wd {
			t.Errorf("Expected path to be '%s', got '%s'", wd, path)
		}
	})
}
