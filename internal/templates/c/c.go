package c

import (
	"path/filepath"

	"github.com/moeenn/projects/internal/templates"
)

func NewCConfig(args *templates.TemplateArgs) *templates.TemplateConfig {
	mainSrc := filepath.Join(args.RootPath, "src")
	binDir := filepath.Join(args.RootPath, "bin")

	files := map[string]string{
		"c.main_c":   filepath.Join(mainSrc, "main.c"),
		"c.makefile": filepath.Join(args.RootPath, "Makefile"),
	}

	return &templates.TemplateConfig{
		Directories: []string{mainSrc, binDir},
		Files:       files,
		Gitignore: []string{
			"bin/*",
			"*.o",
			".DS_Store",
		},
	}
}
