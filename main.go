package main

import (
	"embed"
	"flag"
	"fmt"
	"github.com/moeenn/projects/internal/templates"
	"github.com/moeenn/projects/internal/templates/cpp"
	"github.com/moeenn/projects/internal/templates/javaGradle"
	"github.com/moeenn/projects/internal/templates/js"
	"github.com/moeenn/projects/internal/templates/python"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed stubs
var stubFS embed.FS

func main() {
	templatePtr := flag.String("template", "cpp", "Project template to use. Valid options are 'cpp', 'js', 'java-gradle', 'python'")
	projectNamePtr := flag.String("name", "sandbox", "Name of project being initialized")
	flag.Parse()

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to detect project current directory: %s", err.Error())
		os.Exit(1)
	}

	stubTemplates := template.Must(template.ParseFS(stubFS, "stubs/**/*.stub"))
	templateArgs := &templates.TemplateArgs{
		Templates:   stubTemplates,
		ProjectName: *projectNamePtr,
		RootPath:    filepath.Join(cwd, *projectNamePtr),
	}

	switch *templatePtr {
	case "cpp":
		err = cpp.Initialize(templateArgs)

	case "js":
		err = js.Initialize(templateArgs)

	case "java-gradle":
		err = javaGradle.Initialize(templateArgs)

	case "python":
		err = python.Initialize(templateArgs)

	default:
		err = fmt.Errorf("Invalid project template name: %s", *templatePtr)
	}

	if err != nil {
		// remove any created files in case of error
		os.Remove(templateArgs.RootPath)

		// report the error
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}
}
