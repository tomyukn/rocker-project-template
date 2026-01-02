package generator

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGenerate_FailsIfDirectoryExistsWithoutForce(t *testing.T) {
	tmpDir := t.TempDir()

	projectPath := filepath.Join(tmpDir, "existing")
	if err := os.Mkdir(projectPath, 0o755); err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}

	cfg := ProjectConfig{
		ProjectName: projectPath,
		RVersion:    "4.4",
		ServiceName: "rstudio",
		Force:       false,
	}

	if err := Generate(cfg); err == nil {
		t.Errorf("expected Generate() to fail when directory exists without force")
	}
}
