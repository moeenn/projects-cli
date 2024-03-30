package c

import (
	"fmt"
	"github.com/moeenn/projects/internal/templates"
	"os"
	"path/filepath"
)

func Initialize(args *templates.TemplateArgs) error {
	fmt.Printf("Initializing new C project: %s\n", args.ProjectName)
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
	if err != nil {
		return err
	}

	// create bin directory
	binPath := filepath.Join(args.RootPath, "bin")
	err = os.Mkdir(binPath, os.ModePerm)

	return err
}

func initFiles(args *templates.TemplateArgs) error {
	// create gitignore
	gitignorePath := filepath.Join(args.RootPath, ".gitignore")
	err := templates.CreateFileFromTemplate(args.Templates, gitignorePath, "c.gitignore", nil)
	if err != nil {
		return err
	}

	// create main.cpp file
	mainCPath := filepath.Join(args.RootPath, "src", "main.c")
	err = templates.CreateFileFromTemplate(args.Templates, mainCPath, "c.main_c", nil)
	if err != nil {
		return err
	}

	// create Makefile
	makefilePath := filepath.Join(args.RootPath, "Makefile")
	err = templates.CreateFileFromTemplate(args.Templates, makefilePath, "c.makefile", args)

	return err
}
