package cpp

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

type MakeFileStubArgs struct {
	ProjectName string
}

type CppProjectTemplate struct {
	ProjectName string
	Templates   *template.Template
	cwd         string
}

func NewProject(projectName string, cwd string, tmpl *template.Template) *CppProjectTemplate {
	return &CppProjectTemplate{
		ProjectName: projectName,
		Templates:   tmpl,
		cwd:         cwd,
	}
}

func (pt *CppProjectTemplate) Initialize() error {
	fmt.Printf("Initializing new C++ project: %s\n", pt.ProjectName)
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

func (pt *CppProjectTemplate) initDirectoryStructure() error {
	// create root project folder
	rootPath := filepath.Join(pt.cwd, pt.ProjectName)
	err := os.Mkdir(rootPath, os.ModePerm)
	if err != nil {
		return err
	}

	// create src directory
	srcPath := filepath.Join(rootPath, "src")
	err = os.Mkdir(srcPath, os.ModePerm)
	if err != nil {
		return err
	}

	// create bin directory
	binPath := filepath.Join(rootPath, "bin")
	err = os.Mkdir(binPath, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func (pt *CppProjectTemplate) initFiles() error {
	// create gitignore
	gitignorePath := filepath.Join(pt.cwd, pt.ProjectName, ".gitignore")
	gitignoreFile, err := os.Create(gitignorePath)
	if err != nil {
		return err
	}
	defer gitignoreFile.Close()

	err = pt.Templates.ExecuteTemplate(gitignoreFile, "cpp.gitignore", nil)
	if err != nil {
		return err
	}

	// create main.cpp file
	mainCppPath := filepath.Join(pt.cwd, pt.ProjectName, "src", "main.cpp")
	mainCppFile, err := os.Create(mainCppPath)
	if err != nil {
		return err
	}
	defer mainCppFile.Close()

	err = pt.Templates.ExecuteTemplate(mainCppFile, "cpp.main_cpp", nil)
	if err != nil {
		return err
	}

	// create Makefile
	makefilePath := filepath.Join(pt.cwd, pt.ProjectName, "Makefile")
	makefile, err := os.Create(makefilePath)
	if err != nil {
		return err
	}
	defer makefile.Close()

	makeFileArgs := MakeFileStubArgs{
		ProjectName: pt.ProjectName,
	}

	err = pt.Templates.ExecuteTemplate(makefile, "cpp.makefile", makeFileArgs)
	if err != nil {
		return err
	}

	return nil
}
