package scanners

import (
	"os"
	"path/filepath"
	"testing"

	"Scruticode/internal/shared/constants"
)

const errWritePackageJSON = "failed to write package.json: %v"
const errExpectedStatus = "expected status %d, got %d"

type linterJavascriptExistsCase struct {
	name           string
	packageJSON    string
	expectedStatus int
}

func runLinterJavascriptExistsCase(t *testing.T, tc linterJavascriptExistsCase) {
	t.Helper()
	t.Parallel()

	tmpDir := t.TempDir()
	packageJSONPath := filepath.Join(tmpDir, "package.json")
	if err := os.WriteFile(packageJSONPath, []byte(tc.packageJSON), 0o644); err != nil {
		t.Fatalf(errWritePackageJSON, err)
	}

	response := LinterJavascriptExists(tmpDir)
	if response.Status != tc.expectedStatus {
		t.Errorf(errExpectedStatus, tc.expectedStatus, response.Status)
	}
}

func TestLinterJavascriptExists(t *testing.T) {
	t.Parallel()

	testCases := []linterJavascriptExistsCase{
		{
			name: "valid package.json",
			packageJSON: `{
			"dependencies": {
				"eslint": "8.0.0"
			},
			"scripts": {
				"lint": "eslint ."
			}
		}`,
			expectedStatus: constants.QualityCheckSuccess,
		},
		{
			name: "missing eslint dependency",
			packageJSON: `{
			"dependencies": {},
			"scripts": {
				"lint": "eslint ."
			}
		}`,
			expectedStatus: constants.QualityCheckFailed,
		},
		{
			name: "missing lint script",
			packageJSON: `{
			"dependencies": {
				"eslint": "8.0.0"
			},
			"scripts": {}
		}`,
			expectedStatus: constants.QualityCheckFailed,
		},
		{
			name: "invalid json",
			packageJSON: `{
			"dependencies": {
				"eslint": "8.0.0"
			},
			"scripts": {
				"lint": "eslint ."
			},
		}`,
			expectedStatus: constants.QualityCheckFailed,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			runLinterJavascriptExistsCase(t, tc)
		})
	}
}

func TestPreCommitExists(t *testing.T) {
	t.Parallel()

	t.Run("javascript with husky configured", func(t *testing.T) {
		t.Parallel()
		tempDir := t.TempDir()
		createDir(t, tempDir, ".husky")
		createFile(t, tempDir, "package.json", `{"devDependencies":{"husky":"9.0.0"}}`)

		if PreCommitExists("javascript", tempDir).Status != constants.QualityCheckSuccess {
			t.Fatal("expected PreCommitExists to return true for javascript")
		}
	})

	t.Run("non-javascript with pre-commit file", func(t *testing.T) {
		t.Parallel()
		tempDir := t.TempDir()
		createFile(t, tempDir, ".pre-commit-config.yaml", "repos: []")

		if PreCommitExists("go", tempDir).Status != constants.QualityCheckSuccess {
			t.Fatal("expected PreCommitExists to return true for non-javascript")
		}
	})
}
