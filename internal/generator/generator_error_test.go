package generator

import (
	"os"
	"testing"
)

func TestGenerate_FailsIfDirectoryExists(t *testing.T) {
	tmpDir := t.TempDir()

	projectPath := tmpDir + "/existing"
	if err := os.Mkdir(projectPath, 0o755); err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}

	cfg := ProjectConfig{
		ProjectName: projectPath,
		RVersion:    "4.4",
		ServiceName: "rstudio",
	}

	if err := Generate(cfg); err == nil {
		t.Errorf("expected Generate() to fail when directory exists")
	}
}
