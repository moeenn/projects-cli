package golang

import (
	"path/filepath"

	"github.com/moeenn/projects-cli/internal/templates"
)

// TODO: add a README.md file

func NewGolangConfig(args *templates.TemplateArgs) *templates.TemplateConfig {
	files := map[string]string{
		"go.main_go":      filepath.Join(args.RootPath, "main.go"),
		"go.go_mod":       filepath.Join(args.RootPath, "go.mod"),
		"go.golangci_yml": filepath.Join(args.RootPath, ".golangci.yml"),
	}

	return &templates.TemplateConfig{
		Directories: []string{},
		Files:       files,
		Gitignore: []string{
			args.ProjectName,
			".DS_Store",
		},
	}
}
