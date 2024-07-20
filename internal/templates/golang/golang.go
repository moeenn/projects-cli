package golang

import (
	"path/filepath"

	"github.com/moeenn/projects/internal/templates"
)

func NewGolangConfig(args *templates.TemplateArgs) *templates.TemplateConfig {
	files := map[string]string{
		"go.main_go": filepath.Join(args.RootPath, "main.go"),
		"go.go_mod":  filepath.Join(args.RootPath, "go.mod"),
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
