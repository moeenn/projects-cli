package python

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

type ProjectNameStubArgs struct {
	ProjectName string
}

type PythonProjectTemplate struct {
	ProjectName string
	Templates   *template.Template
	cwd         string
}

func NewProject(projectName string, cwd string, tmpl *template.Template) *PythonProjectTemplate {
	return &PythonProjectTemplate{
		ProjectName: projectName,
		Templates:   tmpl,
		cwd:         cwd,
	}
}

func (pt *PythonProjectTemplate) Initialize() error {
	fmt.Printf("Initializing new Python project: %s\n", pt.ProjectName)
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

func (pt *PythonProjectTemplate) initDirectoryStructure() error {
	// create root project folder
	rootPath := filepath.Join(pt.cwd, pt.ProjectName)
	err := os.Mkdir(rootPath, os.ModePerm)
	if err != nil {
		return err
	}

	// create app directory
	appPath := filepath.Join(rootPath, "app")
	err = os.Mkdir(appPath, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func (pt *PythonProjectTemplate) initFiles() error {
	rootPath := filepath.Join(pt.cwd, pt.ProjectName)
	appPath := filepath.Join(rootPath, "app")

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

	err = pt.Templates.ExecuteTemplate(gitignoreFile, "python.gitignore", nil)
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

	err = pt.Templates.ExecuteTemplate(readmeMdFile, "python.readme_md", projectNameArgs)
	if err != nil {
		return err
	}

	// create __init__py.py file
	initPyPath := filepath.Join(appPath, "__init__.py")
	initPyFile, err := os.Create(initPyPath)
	if err != nil {
		return err
	}
	defer initPyFile.Close()

	err = pt.Templates.ExecuteTemplate(initPyFile, "python.init_py", nil)
	if err != nil {
		return err
	}

	// create app.py file
	appPyPath := filepath.Join(appPath, "app.py")
	appPyFile, err := os.Create(appPyPath)
	if err != nil {
		return err
	}
	defer appPyFile.Close()

	err = pt.Templates.ExecuteTemplate(appPyFile, "python.app_py", nil)
	if err != nil {
		return err
	}

	// create app.py file
	appTestPyPath := filepath.Join(appPath, "app_test.py")
	appTestPyFile, err := os.Create(appTestPyPath)
	if err != nil {
		return err
	}
	defer appTestPyFile.Close()

	err = pt.Templates.ExecuteTemplate(appTestPyFile, "python.app_test_py", nil)
	if err != nil {
		return err
	}

	// create app.py file
	mainPyPath := filepath.Join(rootPath, "main.py")
	mainPyFile, err := os.Create(mainPyPath)
	if err != nil {
		return err
	}
	defer mainPyFile.Close()

	err = pt.Templates.ExecuteTemplate(mainPyFile, "python.main_py", nil)
	if err != nil {
		return err
	}

	// create pyproject.toml file
	pyprojectTomlPath := filepath.Join(rootPath, "pyproject.toml")
	pyprojectTomlFile, err := os.Create(pyprojectTomlPath)
	if err != nil {
		return err
	}
	defer pyprojectTomlFile.Close()

	err = pt.Templates.ExecuteTemplate(pyprojectTomlFile, "python.pyproject_toml", projectNameArgs)
	if err != nil {
		return err
	}

	// create tasks.py file
	tasksPyPath := filepath.Join(rootPath, "tasks.py")
	tasksPyFile, err := os.Create(tasksPyPath)
	if err != nil {
		return err
	}
	defer tasksPyFile.Close()

	err = pt.Templates.ExecuteTemplate(tasksPyFile, "python.tasks_py", nil)
	if err != nil {
		return err
	}

	return nil
}
