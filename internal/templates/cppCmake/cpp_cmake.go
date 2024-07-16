package cppCmake

import (
	"path/filepath"

	"github.com/moeenn/projects/internal/templates"
)

// TODO: add README.md file

func NewCPPCmakeConfig(args *templates.TemplateArgs) *templates.TemplateConfig {
	mainSrc := filepath.Join(args.RootPath, "src")
	files := map[string]string{
		"cpp-cmake.main_cpp":       filepath.Join(mainSrc, "main.cpp"),
		"cpp-cmake.cmakelists_txt": filepath.Join(args.RootPath, "CMakeLists.txt"),
	}

	return &templates.TemplateConfig{
		Directories: []string{mainSrc},
		Files:       files,
		Gitignore: []string{
			"CMakeFiles/*",
			"cmake_install.cmake",
			"CMakeCache.txt",
			"Makefile",
			".cache",
			"compile_commands.json",
			"bin",
			".DS_Store",
		},
	}
}
