package js

import (
	"fmt"
	"github.com/moeenn/projects/internal/templates"
	"os"
	"path/filepath"
)

func Initialize(args *templates.TemplateArgs) error {
	fmt.Printf("Initializing new JS project: %s\n", args.ProjectName)
	err := initDirectoryStructure(args)
	if err != nil {
		return err
	}

	err = initFiles(args)
	return err
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
	srcDirPath := filepath.Join(args.RootPath, "src")

	// create gitignore
	gitignorePath := filepath.Join(args.RootPath, ".gitignore")
	err := templates.CreateFileFromTemplate(args.Templates, gitignorePath, "js.gitignore", args)
	if err != nil {
		return err
	}

	// create README.md file
	readmeMdPath := filepath.Join(args.RootPath, "README.md")
	err = templates.CreateFileFromTemplate(args.Templates, readmeMdPath, "js.readme_md", args)
	if err != nil {
		return err
	}

	// create jsconfig.json file
	jsconfigJsonPath := filepath.Join(args.RootPath, "jsconfig.json")
	err = templates.CreateFileFromTemplate(args.Templates, jsconfigJsonPath, "js.jsconfig_json", args)
	if err != nil {
		return err
	}

	// create package.json file
	packageJsonPath := filepath.Join(args.RootPath, "package.json")
	err = templates.CreateFileFromTemplate(args.Templates, packageJsonPath, "js.package_json", args)
	if err != nil {
		return err
	}

	// create index.mjs file
	indexMjsPath := filepath.Join(srcDirPath, "index.mjs")
	err = templates.CreateFileFromTemplate(args.Templates, indexMjsPath, "js.index_mjs", nil)
	if err != nil {
		return err
	}

	// create index.test.mjs file
	indexTestMjsPath := filepath.Join(srcDirPath, "index.test.mjs")
	err = templates.CreateFileFromTemplate(args.Templates, indexTestMjsPath, "js.index_test_mjs", nil)
	return err
}
