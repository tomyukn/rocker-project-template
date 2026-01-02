package generator

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGenerate_OverwriteWithForce(t *testing.T) {
	tmpDir := t.TempDir()

	projectPath := filepath.Join(tmpDir, "force-project")
	if err := os.Mkdir(projectPath, 0o755); err != nil {
		t.Fatalf("failed to create project dir: %v", err)
	}

	// create a dummy file that should be removed
	dummyFile := filepath.Join(projectPath, "old.txt")
	if err := os.WriteFile(dummyFile, []byte("old"), 0o644); err != nil {
		t.Fatalf("failed to create dummy file: %v", err)
	}

	cfg := ProjectConfig{
		ProjectName: projectPath,
		RVersion:    "4.4",
		ServiceName: "rstudio",
		Force:       true,
	}

	if err := Generate(cfg); err != nil {
		t.Fatalf("Generate() failed with force: %v", err)
	}

	if _, err := os.Stat(dummyFile); !os.IsNotExist(err) {
		t.Errorf("expected old files to be removed when force is enabled")
	}

	if _, err := os.Stat(filepath.Join(projectPath, "Dockerfile")); err != nil {
		t.Errorf("expected Dockerfile to be generated after force overwrite")
	}
}
