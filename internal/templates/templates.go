package templates

import (
	"bytes"
	"os"
	"strings"
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
