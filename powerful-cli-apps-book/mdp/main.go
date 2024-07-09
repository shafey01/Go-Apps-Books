package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

const (
	header = `
      <!DOCTYPE html>
      <html>
        <head>
        <meta http-equiv="content-type" content="text/html; charset=utf-8">
        <title>Markdown Preview Tool</title>
        </head>
        <body>
        `
	footer = `
        </body>
      </html>`
)

func main() {
	// flag for filename
	filename := flag.String("filename", "", "Markdown file")
	flag.Parse()

	// if user did not enter filename exit
	if *filename == "" {
		flag.Usage()
		os.Exit(0)
	}
	if err := run(*filename); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(filename string) error {
	input, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	htmlData := parseContent(input)

	outputFile := fmt.Sprintf("%s.html", filepath.Base(filename))
	fmt.Println(outputFile)

	return saveHTML(outputFile, htmlData)
}

func parseContent(input []byte) []byte {
	output := blackfriday.Run(input)

	body := bluemonday.UGCPolicy().SanitizeBytes(output)

	var buffer bytes.Buffer

	buffer.WriteString(header)
	buffer.Write(body)
	buffer.WriteString(footer)

	return buffer.Bytes()
}

func saveHTML(outputfile string, htmlData []byte) error {
	// r w x
	// 4 2 1
	return os.WriteFile(outputfile, htmlData, 0644)
}
