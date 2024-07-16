package python

import (
	"path/filepath"

	"github.com/moeenn/projects/internal/templates"
)

func NewPythonConfig(args *templates.TemplateArgs) *templates.TemplateConfig {
	appDir := filepath.Join(args.RootPath, "app")
	files := map[string]string{
		"python.app_py":           filepath.Join(appDir, "app.py"),
		"python.app_test_py":      filepath.Join(appDir, "app_test.py"),
		"python.init_py":          filepath.Join(appDir, "__init__.py"),
		"python.main_py":          filepath.Join(args.RootPath, "main.py"),
		"python.readme_md":        filepath.Join(args.RootPath, "README.md"),
		"python.requirements_txt": filepath.Join(args.RootPath, "requirements.txt"),
		"python.tasks_py":         filepath.Join(args.RootPath, "tasks.py"),
	}

	return &templates.TemplateConfig{
		Directories: []string{appDir},
		Files:       files,
		Gitignore: []string{
			".venv",
			"*.pyc",
			"__pycache__",
			".vscode",
			"build/",
			"*.egg-info",
			".*_cache",
			"dist/",
			".DS_Store",
		},
	}
}

// func Initialize(args *templates.TemplateArgs) error {
// 	fmt.Printf("Initializing new Python project: %s\n", args.ProjectName)
// 	err := initDirectoryStructure(args)
// 	if err != nil {
// 		return err
// 	}

// 	err = initFiles(args)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func initDirectoryStructure(args *templates.TemplateArgs) error {
// 	// create root project folder
// 	err := os.Mkdir(args.RootPath, os.ModePerm)
// 	if err != nil {
// 		return err
// 	}

// 	// create app directory
// 	appPath := filepath.Join(args.RootPath, "app")
// 	err = os.Mkdir(appPath, os.ModePerm)

// 	return err
// }

// func initFiles(args *templates.TemplateArgs) error {
// 	appPath := filepath.Join(args.RootPath, "app")

// 	// create gitignore
// 	gitignorePath := filepath.Join(args.RootPath, ".gitignore")
// 	err := templates.CreateFileFromTemplate(args.Templates, gitignorePath, "python.gitignore", nil)
// 	if err != nil {
// 		return err
// 	}

// 	// create README.md file
// 	readmeMdPath := filepath.Join(args.RootPath, "README.md")
// 	err = templates.CreateFileFromTemplate(args.Templates, readmeMdPath, "python.readme_md", args)
// 	if err != nil {
// 		return err
// 	}

// 	// create __init__py.py file
// 	initPyPath := filepath.Join(appPath, "__init__.py")
// 	err = templates.CreateFileFromTemplate(args.Templates, initPyPath, "python.init_py", nil)
// 	if err != nil {
// 		return err
// 	}

// 	// create app.py file
// 	appPyPath := filepath.Join(appPath, "app.py")
// 	err = templates.CreateFileFromTemplate(args.Templates, appPyPath, "python.app_py", nil)
// 	if err != nil {
// 		return err
// 	}

// 	// create app.py file
// 	appTestPyPath := filepath.Join(appPath, "app_test.py")
// 	err = templates.CreateFileFromTemplate(args.Templates, appTestPyPath, "python.app_test_py", nil)
// 	if err != nil {
// 		return err
// 	}

// 	// create app.py file
// 	mainPyPath := filepath.Join(args.RootPath, "main.py")
// 	err = templates.CreateFileFromTemplate(args.Templates, mainPyPath, "python.main_py", nil)
// 	if err != nil {
// 		return err
// 	}

// 	// create requirements.txt file
// 	requirementsTxtPath := filepath.Join(args.RootPath, "requirements.txt")
// 	err = templates.CreateFileFromTemplate(args.Templates, requirementsTxtPath, "python.requirements_txt", nil)
// 	if err != nil {
// 		return err
// 	}

// 	// create tasks.py file
// 	tasksPyPath := filepath.Join(args.RootPath, "tasks.py")
// 	err = templates.CreateFileFromTemplate(args.Templates, tasksPyPath, "python.tasks_py", nil)
// 	return err
// }
