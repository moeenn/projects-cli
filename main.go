package main

import (
	"embed"
	"flag"
	"fmt"
	"github.com/moeenn/projects/internal/cpp"
	"os"
	"text/template"
)

//go:embed stubs
var stubFS embed.FS

func main() {
	templatePtr := flag.String("template", "cpp", "Project template to use. Valid options are 'cpp', 'TODO'")
	projectNamePtr := flag.String("name", "sandbox", "Name of project being initialized")
	flag.Parse()

	cwd, err := os.Getwd()
	if err != nil {
		exit("Failed to detect project current directory: " + err.Error())
	}

	stubTemplates := template.Must(template.ParseFS(stubFS, "stubs/**.stub"))
	cppTemplate := cpp.NewProject(*projectNamePtr, cwd, stubTemplates)

	switch *templatePtr {
	case "cpp":
		err := cppTemplate.Initialize()
		if err != nil {
			exit(err.Error())
		}

	default:
		exit("Invalid project template name: " + *templatePtr)
	}
}

func exit(message string) {
	fmt.Fprintf(os.Stderr, "Error: %s\n", message)
	os.Exit(1)
}
