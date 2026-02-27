package version

import (
	"os/exec"
	"strings"
)

var _version string

func GetVersion() string {
	if _version == "" {
		_version = getVersionFromGit()
	}
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
