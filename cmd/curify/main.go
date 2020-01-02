package main

import (
	"fmt"
	"io"
	"os"

	"github.com/go-shiori/go-readability"
)

func main() {
	var reader io.Reader
	var writer io.Writer = os.Stdout
	var path string

	// Gets the input stream
	args := os.Args[1:]
	if len(args) == 1 {
		// from a file
		file, err := os.Open(args[0])

		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to open '%s': %v", args[0], err)
			os.Exit(1)
		}

		reader = file
		path = args[0]
	} else {
		// from stdin
		reader = os.Stdin
		path = "stdin"
	}

	// Extracts main content from the input stream
	article, err := readability.FromReader(reader, fmt.Sprintf("file://%s", path))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open '%s': %v\n", path, err)
		os.Exit(2)
	}

	writer.Write([]byte(article.Content))
	os.Exit(0)
}
