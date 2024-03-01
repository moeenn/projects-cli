package cppCmake

import (
	"fmt"
	"github.com/moeenn/projects/internal/templates"
	"os"
	"path/filepath"
)

func Initialize(args *templates.TemplateArgs) error {
	fmt.Printf("Initializing new C++ (CMake) project: %s\n", args.ProjectName)
	err := initDirectoryStructure(args)
	if err != nil {
		return err
	}

	err = initFiles(args)
	if err != nil {
		return err
	}

	return nil
}

func initDirectoryStructure(args *templates.TemplateArgs) error {
	// create root project folder
	err := os.Mkdir(args.RootPath, os.ModePerm)
	if err != nil {
		return err
	}

	// create src directory
	srcPath := filepath.Join(args.RootPath, "src")
	err = os.Mkdir(srcPath, os.ModePerm)
	return err
}

func initFiles(args *templates.TemplateArgs) error {
	// create gitignore
	gitignorePath := filepath.Join(args.RootPath, ".gitignore")
	err := templates.CreateFileFromTemplate(args.Templates, gitignorePath, "cpp-cmake.gitignore", nil)
	if err != nil {
		return err
	}

	// create main.cpp file
	mainCppPath := filepath.Join(args.RootPath, "src", "main.cpp")
	err = templates.CreateFileFromTemplate(args.Templates, mainCppPath, "cpp-cmake.main_cpp", nil)
	if err != nil {
		return err
	}

	// create CMakeLists.txt file
	cmakelistsPath := filepath.Join(args.RootPath, "CMakeLists.txt")
	err = templates.CreateFileFromTemplate(args.Templates, cmakelistsPath, "cpp-cmake.cmakelists_txt", args)

	return err
}
