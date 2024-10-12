package php

import (
	"path/filepath"

	"github.com/moeenn/projects/internal/templates"
)

func NewPHPConfig(args *templates.TemplateArgs) *templates.TemplateConfig {
	srcDir := filepath.Join(args.RootPath, "src")
	testsDir := filepath.Join(args.RootPath, "tests")
	files := map[string]string{
		"php.app_php":       filepath.Join(srcDir, "App.php"),
		"php.todo_php":      filepath.Join(srcDir, "Todo.php"),
		"php.todotest_php":  filepath.Join(testsDir, "TodoTest.php"),
		"php.composer_json": filepath.Join(args.RootPath, "composer.json"),
		"php.phpstan_neon":  filepath.Join(args.RootPath, "phpstan.neon"),
		"php.phpunit_xml":   filepath.Join(args.RootPath, "phpunit.xml"),
	}

	return &templates.TemplateConfig{
		Directories: []string{srcDir, testsDir},
		Files:       files,
		Gitignore: []string{
			"cache/*",
			"*.cache",
			"vendor",
		},
	}
}
