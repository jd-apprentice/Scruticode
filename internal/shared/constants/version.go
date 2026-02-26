package constants

import (
	"os"
	"os/exec"
	"strings"
	"sync"
)

var (
	_version     string
	_versionOnce sync.Once
)

func GetVersion() string {
	_versionOnce.Do(func() {
		_version = os.Getenv("APP_VERSION")

		if _version == "" {
			_version = getVersionFromGit()
		}

		if _version == "" {
			_version = "0.0.0-unknown"
		}
	})
	return _version
}

func getVersionFromGit() string {
	cmd := exec.Command("git", "describe", "--tags", "--abbrev=0")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}

	version := strings.TrimSpace(string(output))
	version = strings.TrimPrefix(version, "v")

	return version
}
