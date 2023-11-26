package arguments

import (
	"errors"
	"flag"
)

type Arguments struct {
	Template *string
	Name     *string
}

func Parse() (*Arguments, error) {
	var args Arguments
	args.Template = flag.String("template", "", "Name of the project template to use")
	args.Name = flag.String("name", "", "Name for the new project")

	flag.Parse()

	if *args.Name == "" || *args.Template == "" {
		return &args, errors.New("Missing argument flags. Please see add -h flag for usage information.")
	}

	return &args, nil
}
