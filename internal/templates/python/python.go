package python

import (
	"path/filepath"

	"github.com/moeenn/projects/internal/templates"
)

func NewPythonConfig(args *templates.TemplateArgs) *templates.TemplateConfig {
	appDir := filepath.Join(args.RootPath, "app")
	files := map[string]string{
		"py.app_py":           filepath.Join(appDir, "app.py"),
		"py.app_test_py":      filepath.Join(appDir, "app_test.py"),
		"py.init_py":          filepath.Join(appDir, "__init__.py"),
		"py.main_py":          filepath.Join(args.RootPath, "main.py"),
		"py.readme_md":        filepath.Join(args.RootPath, "README.md"),
		"py.requirements_txt": filepath.Join(args.RootPath, "requirements.txt"),
		"py.tasks_py":         filepath.Join(args.RootPath, "tasks.py"),
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
