package js

import (
	"path/filepath"

	"github.com/moeenn/projects-cli/internal/templates"
)

func NewJSConfig(args *templates.TemplateArgs) *templates.TemplateConfig {
	mainSrc := filepath.Join(args.RootPath, "src")
	files := map[string]string{
		"js.index_mjs":        filepath.Join(mainSrc, "index.mjs"),
		"js.index_test_mjs":   filepath.Join(mainSrc, "index.test.mjs"),
		"js.jsconfig_json":    filepath.Join(args.RootPath, "jsconfig.json"),
		"js.package_json":     filepath.Join(args.RootPath, "package.json"),
		"js.readme_md":        filepath.Join(args.RootPath, "README.md"),
		"js.eslint_config_js": filepath.Join(args.RootPath, "eslint.config.js"),
	}

	return &templates.TemplateConfig{
		Directories: []string{mainSrc},
		Files:       files,
		Gitignore: []string{
			"node_modules/",
			".DS_Store",
		},
	}
}
