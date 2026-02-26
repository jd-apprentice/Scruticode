package constants

import (
	"os/exec"
	"strings"
	"sync"
)

var (
	_version     string
	_versionOnce sync.Once
)

// GetVersion returns the version from git tags
func GetVersion() string {
	_versionOnce.Do(func() {
		_version = getVersionFromGit()
	})
	return _version
}

func getVersionFromGit() string {
	cmd := exec.Command("git", "describe", "--tags", "--abbrev=0")
	output, _ := cmd.Output()

	version := strings.TrimSpace(string(output))
	version = strings.TrimPrefix(version, "v")

	return version
}
