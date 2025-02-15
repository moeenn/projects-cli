package main

import (
	"embed"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/moeenn/projects-cli/internal/templates"
	"github.com/moeenn/projects-cli/internal/templates/c"
	"github.com/moeenn/projects-cli/internal/templates/cppCmake"
	"github.com/moeenn/projects-cli/internal/templates/cppMake"
	"github.com/moeenn/projects-cli/internal/templates/golang"
	"github.com/moeenn/projects-cli/internal/templates/javaGradle"
	"github.com/moeenn/projects-cli/internal/templates/js"
	"github.com/moeenn/projects-cli/internal/templates/php"
	"github.com/moeenn/projects-cli/internal/templates/python"
	"github.com/moeenn/projects-cli/internal/templates/ts"
)

//go:embed stubs
var stubFS embed.FS

var (
	TEMPLATE_NAMES = [9]string{
		"go", "c", "cpp-cmake", "cpp-make", "javascript (or 'js')", "typescript (or 'ts')", "java-gradle", "python (or 'py')", "php",
	}
)

func run() error {
	templatePtr := flag.String("template", TEMPLATE_NAMES[0], "Project template to use")
	projectNamePtr := flag.String("name", "sandbox", "Name of project being initialized")
	listTemplatesPtr := flag.Bool("list", false, "Print list of available template names")
	flag.Parse()

	if *listTemplatesPtr {
		fmt.Printf("Valid templates include: \n - %s\n", strings.Join(TEMPLATE_NAMES[:], "\n - "))
		return nil
	}

	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to detect project current directory: %s", err.Error())
	}

	stubTemplates := template.Must(template.ParseFS(stubFS, "stubs/**/*.stub"))
	templateArgs := &templates.TemplateArgs{
		Templates:   stubTemplates,
		ProjectName: *projectNamePtr,
		RootPath:    filepath.Join(cwd, *projectNamePtr),
	}

	var config *templates.TemplateConfig
	switch *templatePtr {
	case "c":
		config = c.NewCConfig(templateArgs)
		err = templateArgs.Initialize("C", config)

	case "cpp-make":
		config = cppMake.NewCPPCmakeConfig(templateArgs)
		err = templateArgs.Initialize("C++ (Make)", config)

	case "cpp-cmake":
		config = cppCmake.NewCPPCmakeConfig(templateArgs)
		err = templateArgs.Initialize("C++ (Cmake)", config)

	case "go":
		config = golang.NewGolangConfig(templateArgs)
		err = templateArgs.Initialize("Golang", config)

	case "js", "javascript":
		config = js.NewJSConfig(templateArgs)
		err = templateArgs.Initialize("Javascript (vanilla)", config)

	case "ts", "typescript":
		config = ts.NewTSConfig(templateArgs)
		err = templateArgs.Initialize("Typescript", config)

	case "java-gradle":
		config = javaGradle.NewJavaGradleConfig(templateArgs)
		err = templateArgs.Initialize("Java (Gradle)", config)

	case "py", "python":
		config = python.NewPythonConfig(templateArgs)
		err = templateArgs.Initialize("Python", config)

	case "php":
		config = php.NewPHPConfig(templateArgs)
		err = templateArgs.Initialize("PHP", config)

	default:
		err = fmt.Errorf("invalid project template name: %s", *templatePtr)
	}

	if err != nil {
		// cleanup: remove any created files in case of error
		err = os.RemoveAll(templateArgs.RootPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		}

		return err
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}
}
