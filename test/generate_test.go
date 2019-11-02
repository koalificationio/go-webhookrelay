//nolint:gosec,dogsled
package test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

const (
	openAPIPackage = "pkg/openapi"
	testPackage    = "test_pkg"
)

func TestMain(m *testing.M) {
	tmpDir := filepath.Join(getRoot(), testPackage)

	os.RemoveAll(tmpDir)

	err := os.Mkdir(tmpDir, 0700)
	if err != nil {
		fmt.Printf("Failed creating test folder: %v", err)
		os.Exit(1)
	}

	exitCode := m.Run()

	os.RemoveAll(tmpDir)

	os.Exit(exitCode)
}

func TestOpenAPIGeneration(t *testing.T) {
	root := getRoot()
	generate := exec.Command("make", "generate", fmt.Sprintf("OPENAPI_PACCKAGE=./%s", testPackage))
	generate.Dir = root

	output, err := generate.CombinedOutput()
	if err != nil {
		t.Log(root)
		t.Log(string(output))
		t.Fatalf("Failed running Openapi Generation: %v", err)
	}

	args := []string{
		filepath.Join(root, testPackage),
		"-type",
		"f",
		"-exec",
		"sed",
		"-i",
		"-e",
		"s/test_pkg/pkg\\/openapi/g",
		"{}",
		"+",
	}

	// disable sed file backups on MacOS, still less code than doing this in go
	if runtime.GOOS == "darwin" {
		// insert '' after -i
		args = append(args[:6], append([]string{""}, args[6:]...)...)
	}

	sed := exec.Command("find", args...)
	sed.Dir = root

	output, err = sed.CombinedOutput()
	if err != nil {
		t.Log(string(output))
		t.Log(args)
		t.Fatalf("Failed diffing generated files: %v", err)
	}

	args = []string{
		"diff",
		"--no-index",
		filepath.Join(root, openAPIPackage),
		filepath.Join(root, testPackage),
	}
	diff := exec.Command("git", args...)
	diff.Dir = root

	output, _ = diff.CombinedOutput()
	if d := string(output); d != "" {
		t.Errorf("Mismatch in generated OpenAPI package:\n%s", d)
	}
}

func getRoot() string {
	_, b, _, _ := runtime.Caller(0)

	// Root folder of this project
	root := filepath.Join(filepath.Dir(b), "..")

	return root
}
