package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/lunny/html2md"
)

func main() {
	var reader io.Reader
	var writer io.Writer = os.Stdout

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
	} else {
		// from stdin
		reader = os.Stdin
	}

	// Copies the content
	var content bytes.Buffer
	_, err := io.Copy(&content, reader)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read the content: %v", err)
		os.Exit(2)
	}

	// Converts the content into markdown
	md := html2md.Convert(content.String())
	writer.Write([]byte(md))
	os.Exit(0)
}
