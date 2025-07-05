package functions

import (
	"Scruticode/src/shared/constants"
	"log"
	"os"
	"testing"
)

const errMessage = "Expected %v, got %v"

func setupSuite(tb testing.TB) func(tb testing.TB) {
	tb.Helper()
	log.Println("Setup suite", tb.Name())

	return func(_ testing.TB) {
		os.Remove("Dockerfile")
		os.RemoveAll("docker")
		os.RemoveAll("infra")
	}
}

func createTempDockerfile(t *testing.T, path string) {
	t.Helper()

	if _, err := os.Stat(path); os.IsExist(err) {
		removeTempDockerfile(t, path)
	}

	if err := os.WriteFile(path, []byte("FROM alpine"), 0600); err != nil {
		t.Fatalf("Failed to create temporary Dockerfile: %v", err)
	}
}

func removeTempDockerfile(t *testing.T, path string) {
	t.Helper()
	if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
		t.Fatalf("Failed to remove temporary Dockerfile: %v", err)
	}
}

func removeTempFolder(t *testing.T, path string) {
	t.Helper()
	if err := os.RemoveAll(path); err != nil && !os.IsNotExist(err) {
		t.Fatalf("Failed to remove temporary folder: %v", err)
	}
}

func TestDockerfileExistsInRoot(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	path := "Dockerfile"
	createTempDockerfile(t, path)
	defer removeTempDockerfile(t, path)

	result := DockerfileExists(constants.CurrentPath)
	if result.Status != constants.QualityCheckSuccess {
		t.Errorf(errMessage, constants.QualityCheckSuccess, result.Status)
	}
}

func TestDockerfileExistsInDockerFolder(t *testing.T) {
	const dockerFolder = "docker"

	const infraFolder = "infra"

	errDockerFolder := os.Mkdir(dockerFolder, 0755)
	errInfraFolder := os.Mkdir(infraFolder, 0755)

	if errDockerFolder != nil {
		t.Fatalf("Failed to create temporary folder: %v", errDockerFolder)
	}

	if errInfraFolder != nil {
		t.Fatalf("Failed to create temporary folder: %v", errInfraFolder)
	}

	createTempDockerfile(t, dockerFolder+"/Dockerfile")
	defer removeTempDockerfile(t, dockerFolder+"/Dockerfile")
	defer removeTempFolder(t, dockerFolder)
	defer removeTempFolder(t, infraFolder)

	result := DockerfileExists("docker")
	if result.Status != constants.QualityCheckSuccess {
		t.Errorf(errMessage, constants.QualityCheckSuccess, result.Status)
	}
}
