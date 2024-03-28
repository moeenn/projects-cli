package main

import (
	"embed"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/moeenn/projects/internal/templates"
	"github.com/moeenn/projects/internal/templates/cppCmake"
	"github.com/moeenn/projects/internal/templates/cppMake"
	"github.com/moeenn/projects/internal/templates/javaGradle"
	"github.com/moeenn/projects/internal/templates/js"
	"github.com/moeenn/projects/internal/templates/python"
)

//go:embed stubs
var stubFS embed.FS

var (
	TEMPLATE_NAMES = [5]string{
		"cpp-cmake", "cpp-make", "javascript (or 'js')", "java-gradle", "python",
	}
)

func main() {
	templatePtr := flag.String("template", TEMPLATE_NAMES[0], "Project template to use")
	projectNamePtr := flag.String("name", "sandbox", "Name of project being initialized")
	listTemplatesPtr := flag.Bool("list", false, "Print list of available template names")
	flag.Parse()

	if *listTemplatesPtr {
		fmt.Printf("Valid templates include: \n - %s\n", strings.Join(TEMPLATE_NAMES[:], "\n - "))
		return
	}

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
	case "cpp-make":
		err = cppMake.Initialize(templateArgs)

	case "cpp-cmake":
		err = cppCmake.Initialize(templateArgs)

	case "js":
		err = js.Initialize(templateArgs)

	case "java-gradle":
		err = javaGradle.Initialize(templateArgs)

	case "python":
		err = python.Initialize(templateArgs)

	default:
		err = fmt.Errorf("invalid project template name: %s", *templatePtr)
	}

	if err != nil {
		// remove any created files in case of error
		_ = os.Remove(templateArgs.RootPath)

		// report the error
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}
}
