package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

type ProjectConfig struct {
	ProjectName string
	RVersion    string
	ServiceName string
	Force       bool
}

func Generate(cfg ProjectConfig) error {
	if _, err := os.Stat(cfg.ProjectName); err == nil {
		if !cfg.Force {
			return fmt.Errorf("project directory already exists: %s", cfg.ProjectName)
		}
		if err := os.RemoveAll(cfg.ProjectName); err != nil {
			return fmt.Errorf("failed to remove existing directory: %w", err)
		}
	}

	if err := os.Mkdir(cfg.ProjectName, 0o755); err != nil {
		return fmt.Errorf("failed to create project directory: %w", err)
	}

	files := map[string]string{
		"Dockerfile":   "templates/Dockerfile.tmpl",
		"compose.yaml": "templates/compose.yaml.tmpl",
		"README.md":    "templates/README.md.tmpl",
	}

	for output, tmplPath := range files {
		if err := renderTemplate(
			filepath.Join(cfg.ProjectName, output),
			tmplPath,
			cfg,
		); err != nil {
			return err
		}
	}

	return nil
}

func renderTemplate(outputPath, templatePath string, data any) error {
	tmpl, err := template.ParseFS(templateFS, templatePath)
	if err != nil {
		return fmt.Errorf("failed to parse embedded template %s: %w", templatePath, err)
	}

	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", outputPath, err)
	}
	defer f.Close()

	if err := tmpl.Execute(f, data); err != nil {
		return fmt.Errorf("failed to render template %s: %w", templatePath, err)
	}

	return nil
}
