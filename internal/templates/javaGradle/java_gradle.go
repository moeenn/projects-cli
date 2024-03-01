package javaGradle

import (
	"fmt"
	"github.com/moeenn/projects/internal/templates"
	"os"
	"path/filepath"
)

func Initialize(args *templates.TemplateArgs) error {
	fmt.Printf("Initializing new Java-Gradle project: %s\n", args.ProjectName)
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
	srcPath := filepath.Join(args.RootPath, "src", "main", "java", "com", args.ProjectName)
	err = os.MkdirAll(srcPath, os.ModePerm)
	if err != nil {
		return err
	}

	// create test directory
	testPath := filepath.Join(args.RootPath, "src", "test", "java", "com", args.ProjectName)
	err = os.MkdirAll(testPath, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func initFiles(args *templates.TemplateArgs) error {
	srcPath := filepath.Join(args.RootPath, "src", "main", "java", "com", args.ProjectName)
	testPath := filepath.Join(args.RootPath, "src", "test", "java", "com", args.ProjectName)

	// create gitignore
	gitignorePath := filepath.Join(args.RootPath, ".gitignore")
	err := templates.CreateFileFromTemplate(args.Templates, gitignorePath, "java-gradle.gitignore", nil)
	if err != nil {
		return err
	}

	// create README.md file
	readmeMdPath := filepath.Join(args.RootPath, "README.md")
	err = templates.CreateFileFromTemplate(args.Templates, readmeMdPath, "java-gradle.readme_md", args)
	if err != nil {
		return err
	}

	// create build.gradle file
	buildGradlePath := filepath.Join(args.RootPath, "build.gradle")
	err = templates.CreateFileFromTemplate(args.Templates, buildGradlePath, "java-gradle.build_gradle", args)
	if err != nil {
		return err
	}

	// create Main.java file
	mainJavaPath := filepath.Join(srcPath, "Main.java")
	err = templates.CreateFileFromTemplate(args.Templates, mainJavaPath, "java-gradle.main_java", args)
	if err != nil {
		return err
	}

	// create Main.java file
	mainTestJavaPath := filepath.Join(testPath, "MainTest.java")
	err = templates.CreateFileFromTemplate(args.Templates, mainTestJavaPath, "java-gradle.main_test_java", args)
	return err
}
