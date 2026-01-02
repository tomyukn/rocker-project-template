package cmd

import (
	"os"
	"path/filepath"
	"testing"
)

func TestInitCommand(t *testing.T) {
	tmpDir := t.TempDir()
	oldWd, _ := os.Getwd()
	defer os.Chdir(oldWd)

	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("failed to change dir: %v", err)
	}

	cmd := initCmd()

	cmd.SetArgs([]string{"test-project", "--r-version", "4.4"})
	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)

	if err := cmd.Execute(); err != nil {
		t.Fatalf("init command failed: %v", err)
	}

	projectPath := filepath.Join(tmpDir, "test-project")
	if _, err := os.Stat(projectPath); err != nil {
		t.Errorf("expected project directory to be created")
	}
}
