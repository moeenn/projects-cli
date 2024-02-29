package js

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

type ProjectNameStubArgs struct {
	ProjectName string
}

type JsProjectTemplate struct {
	ProjectName string
	Templates   *template.Template
	cwd         string
}

func NewProject(projectName string, cwd string, tmpl *template.Template) *JsProjectTemplate {
	return &JsProjectTemplate{
		ProjectName: projectName,
		Templates:   tmpl,
		cwd:         cwd,
	}
}

func (pt *JsProjectTemplate) Initialize() error {
	fmt.Printf("Initializing new JS project: %s\n", pt.ProjectName)
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

func (pt *JsProjectTemplate) initDirectoryStructure() error {
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

	return nil
}

func (pt *JsProjectTemplate) initFiles() error {
	rootPath := filepath.Join(pt.cwd, pt.ProjectName)
	srcDirPath := filepath.Join(rootPath, "src")

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

	err = pt.Templates.ExecuteTemplate(gitignoreFile, "js.gitignore", nil)
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

	err = pt.Templates.ExecuteTemplate(readmeMdFile, "js.readme_md", projectNameArgs)
	if err != nil {
		return err
	}

	// create jsconfig.json file
	jsconfigJsonPath := filepath.Join(rootPath, "jsconfig.json")
	jsconfigJsonFile, err := os.Create(jsconfigJsonPath)
	if err != nil {
		return err
	}
	defer jsconfigJsonFile.Close()

	err = pt.Templates.ExecuteTemplate(jsconfigJsonFile, "js.jsconfig_json", projectNameArgs)
	if err != nil {
		return err
	}

	// create package.json file
	packageJsonPath := filepath.Join(rootPath, "package.json")
	packageJsonFile, err := os.Create(packageJsonPath)
	if err != nil {
		return err
	}
	defer packageJsonFile.Close()

	err = pt.Templates.ExecuteTemplate(packageJsonFile, "js.package_json", projectNameArgs)
	if err != nil {
		return err
	}

	// create index.mjs file
	indexMjsPath := filepath.Join(srcDirPath, "index.mjs")
	indexMjsFile, err := os.Create(indexMjsPath)
	if err != nil {
		return err
	}
	defer indexMjsFile.Close()

	err = pt.Templates.ExecuteTemplate(indexMjsFile, "js.index_mjs", nil)
	if err != nil {
		return err
	}

	// create index.test.mjs file
	indexTestMjsPath := filepath.Join(srcDirPath, "index.test.mjs")
	indexTestMjsFile, err := os.Create(indexTestMjsPath)
	if err != nil {
		return err
	}
	defer indexTestMjsFile.Close()

	err = pt.Templates.ExecuteTemplate(indexTestMjsFile, "js.index_test_mjs", nil)
	if err != nil {
		return err
	}

	// create entity.mjs file
	entityMjsPath := filepath.Join(srcDirPath, "entity.mjs")
	entityMjsFile, err := os.Create(entityMjsPath)
	if err != nil {
		return err
	}
	defer entityMjsFile.Close()

	err = pt.Templates.ExecuteTemplate(entityMjsFile, "js.entity_mjs", nil)
	if err != nil {
		return err
	}

	return nil
}
