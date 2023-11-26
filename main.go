package main

import (
	"fmt"
	"os"
	"projects/pkg/arguments"
)

func main() {
	args, err := arguments.Parse()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err.Error())
		return
	}

	fmt.Printf("%v, %v\n", *args.Name, *args.Template)
}
