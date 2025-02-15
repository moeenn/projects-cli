package ts

import (
	"path/filepath"

	"github.com/moeenn/projects-cli/internal/templates"
)

func NewTSConfig(args *templates.TemplateArgs) *templates.TemplateConfig {
	mainSrc := filepath.Join(args.RootPath, "src")
	files := map[string]string{
		"ts.index_test_ts": filepath.Join(mainSrc, "index.test.ts"),
		"ts.index_ts":      filepath.Join(mainSrc, "index.ts"),
		"ts.package_json":  filepath.Join(args.RootPath, "package.json"),
		"ts.readme_md":     filepath.Join(args.RootPath, "README.md"),
		"ts.swrrc":         filepath.Join(args.RootPath, ".swrrc"),
		"ts.tsconfig_json": filepath.Join(args.RootPath, "tsconfig.json"),
	}

	return &templates.TemplateConfig{
		Directories: []string{mainSrc},
		Files:       files,
		Gitignore: []string{
			"node_modules",
			"build",
			".DS_Store",
		},
	}
}

// func Initialize(args *templates.TemplateArgs) error {
// 	fmt.Printf("Initializing new TS project: %s\n", args.ProjectName)
// 	err := initDirectoryStructure(args)
// 	if err != nil {
// 		return err
// 	}

// 	err = initFiles(args)
// 	return err
// }

// func initDirectoryStructure(args *templates.TemplateArgs) error {
// 	// create root project folder
// 	err := os.Mkdir(args.RootPath, os.ModePerm)
// 	if err != nil {
// 		return err
// 	}

// 	// create src directory
// 	srcPath := filepath.Join(args.RootPath, "src")
// 	err = os.Mkdir(srcPath, os.ModePerm)

// 	return err
// }

// func initFiles(args *templates.TemplateArgs) error {
// 	srcDirPath := filepath.Join(args.RootPath, "src")

// 	// create gitignore
// 	gitignorePath := filepath.Join(args.RootPath, ".gitignore")
// 	err := templates.CreateFileFromTemplate(args.Templates, gitignorePath, "ts.gitignore", args)
// 	if err != nil {
// 		return err
// 	}

// 	// create README.md file
// 	readmeMdPath := filepath.Join(args.RootPath, "README.md")
// 	err = templates.CreateFileFromTemplate(args.Templates, readmeMdPath, "ts.readme_md", args)
// 	if err != nil {
// 		return err
// 	}

// 	// create tsconfig.json file
// 	tsconfigJsonPath := filepath.Join(args.RootPath, "tsconfig.json")
// 	err = templates.CreateFileFromTemplate(args.Templates, tsconfigJsonPath, "ts.tsconfig_json", args)
// 	if err != nil {
// 		return err
// 	}

// 	// create .swrrc file
// 	swrrcPath := filepath.Join(args.RootPath, ".swrrc")
// 	err = templates.CreateFileFromTemplate(args.Templates, swrrcPath, "ts.swrrc", args)
// 	if err != nil {
// 		return err
// 	}

// 	// create package.json file
// 	packageJsonPath := filepath.Join(args.RootPath, "package.json")
// 	err = templates.CreateFileFromTemplate(args.Templates, packageJsonPath, "ts.package_json", args)
// 	if err != nil {
// 		return err
// 	}

// 	// create index.ts file
// 	indexTsPath := filepath.Join(srcDirPath, "index.ts")
// 	err = templates.CreateFileFromTemplate(args.Templates, indexTsPath, "ts.index_ts", nil)
// 	if err != nil {
// 		return err
// 	}

// 	// create index.test.ts file
// 	indexTestTsPath := filepath.Join(srcDirPath, "index.test.ts")
// 	err = templates.CreateFileFromTemplate(args.Templates, indexTestTsPath, "ts.index_test_ts", nil)
// 	return err
// }
