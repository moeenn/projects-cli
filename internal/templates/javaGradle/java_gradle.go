package javaGradle

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

type ProjectNameStubArgs struct {
	ProjectName string
}

type JavaGradleProjectTemplate struct {
	ProjectName string
	Templates   *template.Template
	cwd         string
}

func NewProject(projectName string, cwd string, tmpl *template.Template) *JavaGradleProjectTemplate {
	return &JavaGradleProjectTemplate{
		ProjectName: projectName,
		Templates:   tmpl,
		cwd:         cwd,
	}
}

func (pt *JavaGradleProjectTemplate) Initialize() error {
	fmt.Printf("Initializing new Java-Gradle project: %s\n", pt.ProjectName)
	err := pt.initDirectoryStructure()
	if err != nil {
		return err
	}

	err = pt.initFiles()
	if err != nil {
		return err
	}

	return nil
}

func (pt *JavaGradleProjectTemplate) initDirectoryStructure() error {
	// create root project folder
	rootPath := filepath.Join(pt.cwd, pt.ProjectName)
	err := os.Mkdir(rootPath, os.ModePerm)
	if err != nil {
		return err
	}

	// create src directory
	srcPath := filepath.Join(rootPath, "src", "main", "java", "com", pt.ProjectName)
	err = os.MkdirAll(srcPath, os.ModePerm)
	if err != nil {
		return err
	}

	// create test directory
	testPath := filepath.Join(rootPath, "src", "test", "java", "com", pt.ProjectName)
	err = os.MkdirAll(testPath, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func (pt *JavaGradleProjectTemplate) initFiles() error {
	rootPath := filepath.Join(pt.cwd, pt.ProjectName)
	srcPath := filepath.Join(rootPath, "src", "main", "java", "com", pt.ProjectName)
	testPath := filepath.Join(rootPath, "src", "test", "java", "com", pt.ProjectName)

	projectNameArgs := ProjectNameStubArgs{
		ProjectName: pt.ProjectName,
	}

	// create gitignore
	gitignorePath := filepath.Join(rootPath, ".gitignore")
	gitignoreFile, err := os.Create(gitignorePath)
	if err != nil {
		return err
	}
	defer gitignoreFile.Close()

	err = pt.Templates.ExecuteTemplate(gitignoreFile, "java-gradle.gitignore", nil)
	if err != nil {
		return err
	}

	// create README.md file
	readmeMdPath := filepath.Join(rootPath, "README.md")
	readmeMdFile, err := os.Create(readmeMdPath)
	if err != nil {
		return err
	}
	defer readmeMdFile.Close()

	err = pt.Templates.ExecuteTemplate(readmeMdFile, "java-gradle.readme_md", projectNameArgs)
	if err != nil {
		return err
	}

	// create build.gradle file
	buildGradlePath := filepath.Join(rootPath, "build.gradle")
	buildGradleFile, err := os.Create(buildGradlePath)
	if err != nil {
		return err
	}
	defer buildGradleFile.Close()

	err = pt.Templates.ExecuteTemplate(buildGradleFile, "java-gradle.build_gradle", projectNameArgs)
	if err != nil {
		return err
	}

	// create Main.java file
	mainJavaPath := filepath.Join(srcPath, "Main.java")
	mainJavaFile, err := os.Create(mainJavaPath)
	if err != nil {
		return err
	}
	defer mainJavaFile.Close()

	err = pt.Templates.ExecuteTemplate(mainJavaFile, "java-gradle.main_java", projectNameArgs)
	if err != nil {
		return err
	}

	// create Main.java file
	mainTestJavaPath := filepath.Join(testPath, "MainTest.java")
	mainTestJavaFile, err := os.Create(mainTestJavaPath)
	if err != nil {
		return err
	}
	defer mainTestJavaFile.Close()

	err = pt.Templates.ExecuteTemplate(mainTestJavaFile, "java-gradle.main_test_java", projectNameArgs)
	if err != nil {
		return err
	}

	return nil
}
