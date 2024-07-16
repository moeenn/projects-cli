package cppMake

import (
	"path/filepath"

	"github.com/moeenn/projects/internal/templates"
)

// TODO: add README.md file

func NewCPPCmakeConfig(args *templates.TemplateArgs) *templates.TemplateConfig {
	mainSrc := filepath.Join(args.RootPath, "src")
	binDir := filepath.Join(args.RootPath, "bin")

	files := map[string]string{
		"cpp-make.main_cpp": filepath.Join(mainSrc, "main.cpp"),
		"cpp-make.makefile": filepath.Join(args.RootPath, "Makefile"),
	}

	return &templates.TemplateConfig{
		Directories: []string{mainSrc, binDir},
		Files:       files,
		Gitignore: []string{
			"bin/*",
			"*.o",
			".cache",
			"compile_commands.json",
			".DS_Store",
		},
	}
}
