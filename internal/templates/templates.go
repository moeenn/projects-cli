package templates

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"
)

type TemplateArgs struct {
	ProjectName string
	Templates   *template.Template
	RootPath    string
}

func (a *TemplateArgs) Initialize(key string, config *TemplateConfig) error {
	fmt.Printf("Initializing new %s project: %s\n", key, a.ProjectName)

	// create main project folder
	err := os.MkdirAll(a.RootPath, os.ModePerm)
	if err != nil {
		return err
	}

	if err := config.CreateDirectories(); err != nil {
		return err
	}

	if err := config.RenderFiles(a); err != nil {
		return err
	}

	// TODO: generate gitignore file

	return nil
}

func CreateFileFromTemplate(templates *template.Template, filePath string, templateName string, data any) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()

	buf := &bytes.Buffer{}
	err = templates.ExecuteTemplate(buf, templateName, data)
	if err != nil {
		return err
	}

	// templates can be executed directly to files (because file implements Writer
	// interface). However this results in extra whitespace at the start and end
	// of the file. Following is a fix to prevent that.
	asString := buf.String()
	_, err = file.WriteString(strings.TrimSpace(asString))
	if err != nil {
		return err
	}

	return nil
}

type TemplateConfig struct {
	Directories []string
	Files       map[string]string
	Gitignore   []string
}

func (c *TemplateConfig) CreateDirectories() error {
	for _, directory := range c.Directories {
		err := os.MkdirAll(directory, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *TemplateConfig) RenderFiles(args *TemplateArgs) error {
	for key, value := range c.Files {
		err := CreateFileFromTemplate(args.Templates, value, key, args)
		if err != nil {
			return err
		}
	}

	return nil
}

// TODO: implement method on TemplateConfig to render gitignore file
