package javaGradle

import (
	"path/filepath"

	"github.com/moeenn/projects/internal/templates"
)

// TODO: implement similar function for the rest of the templates
func NewJavaGradleConfig(args *templates.TemplateArgs) *templates.TemplateConfig {
	mainSrc := filepath.Join(args.RootPath, "src", "main", "java", args.ProjectName)
	testSrc := filepath.Join(args.RootPath, "src", "test", "java", args.ProjectName)

	files := map[string]string{
		"java-gradle.readme_md":      filepath.Join(args.RootPath, "README.md"),
		"java-gradle.build_gradle":   filepath.Join(args.RootPath, "build.gradle"),
		"java-gradle.main_java":      filepath.Join(mainSrc, "Main.java"),
		"java-gradle.main_test_java": filepath.Join(testSrc, "MainTest.java"),
	}

	return &templates.TemplateConfig{
		Directories: []string{mainSrc, testSrc},
		Files:       files,
		Gitignore: []string{
			".gradle",
			".vscode",
			"bin",
			"build",
			"*.class",
			".DS_Store",
		},
	}
}
