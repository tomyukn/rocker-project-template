package generator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerate_CreatesProjectStructure(t *testing.T) {
	tmpDir := t.TempDir()

	cfg := ProjectConfig{
		ProjectName: filepath.Join(tmpDir, "test-project"),
		RVersion:    "4.4",
		ServiceName: "rstudio",
	}

	if err := Generate(cfg); err != nil {
		t.Fatalf("Generate() failed: %v", err)
	}

	expectedFiles := []string{
		"Dockerfile",
		"compose.yaml",
		"README.md",
	}

	for _, f := range expectedFiles {
		path := filepath.Join(cfg.ProjectName, f)
		if _, err := os.Stat(path); err != nil {
			t.Errorf("expected file %s to exist, but it does not", path)
		}
	}
}

func TestGenerate_DockerfileContent(t *testing.T) {
	tmpDir := t.TempDir()

	cfg := ProjectConfig{
		ProjectName: filepath.Join(tmpDir, "content-test"),
		RVersion:    "4.4",
		ServiceName: "rstudio",
	}

	if err := Generate(cfg); err != nil {
		t.Fatalf("Generate() failed: %v", err)
	}

	data, err := os.ReadFile(filepath.Join(cfg.ProjectName, "Dockerfile"))
	if err != nil {
		t.Fatalf("failed to read Dockerfile: %v", err)
	}

	if !strings.Contains(string(data), "FROM rocker/rstudio:4.4") {
		t.Errorf("Dockerfile does not contain expected base image")
	}
}
