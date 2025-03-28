package cppCmake

import (
	"path/filepath"

	"github.com/moeenn/projects-cli/internal/templates"
)

func NewCPPCmakeConfig(args *templates.TemplateArgs) *templates.TemplateConfig {
	mainSrc := filepath.Join(args.RootPath, "src")
	incluesDir := filepath.Join(args.RootPath, "include")
	buildDir := filepath.Join(args.RootPath, "build")

	files := map[string]string{
		"cpp-cmake.readme_md":      filepath.Join(args.RootPath, "README.md"),
		"cpp-cmake.main_cpp":       filepath.Join(mainSrc, "main.cpp"),
		"cpp-cmake.cmakelists_txt": filepath.Join(args.RootPath, "CMakeLists.txt"),
	}

	return &templates.TemplateConfig{
		Directories: []string{mainSrc, incluesDir, buildDir},
		Files:       files,
		Gitignore: []string{
			".cache",
			"compile_commands.json",
			"build",
			".DS_Store",
		},
	}
}
