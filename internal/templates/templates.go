package templates

import (
	"os"
	"text/template"
)

type TemplateArgs struct {
	ProjectName string
	Templates   *template.Template
	RootPath    string
}

func CreateFileFromTemplate(templates *template.Template, filePath string, templateName string, data any) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()

	err = templates.ExecuteTemplate(file, templateName, data)
	if err != nil {
		return err
	}

	return nil
}
