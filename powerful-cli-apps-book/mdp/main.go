package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io"
	"os"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

// const (
// 	header = `
//       <!DOCTYPE html>
//       <html>
//         <head>
//         <meta http-equiv="content-type" content="text/html; charset=utf-8">
//         <title>Markdown Preview Tool</title>
//         </head>
//         <body>
//         `
// 	footer = `
//         </body>
//       </html>`
// )

const (
	defaultTemplate = `<!DOCTYPE html>
	<html>
	<head>
	<meta http-equiv="content-type" content="text/html; charset=utf-8">
	<title>{{ .Title }}</title>
	</head>
	<body>
	{{ .Body }}
	</body>
	</html>
	`
)

type content struct {
	Title string
	Body  template.HTML
}

func main() {
	// flag for filename
	filename := flag.String("filename", "", "Markdown file")
	fname := flag.String("fname", "", "Alternative file")
	flag.Parse()

	// if user did not enter filename exit
	if *filename == "" {
		flag.Usage()
		os.Exit(0)
	}
	if err := run(*filename, *fname, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(filename string, fname string, out io.Writer) error {
	input, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	htmlData, err := parseContent(input, fname)

	if err != nil {
		return err
	}
	// outputFile := fmt.Sprintf("%s.html", filepath.Base(filename))
	temp, err := os.CreateTemp("", "mdp*.html")

	if err != nil {
		return err
	}

	if err := temp.Close(); err != nil {
		return err
	}

	outputFile := temp.Name()

	fmt.Fprintln(out, outputFile)

	return saveHTML(outputFile, htmlData)
}

func parseContent(input []byte, fname string) ([]byte, error) {
	output := blackfriday.Run(input)

	body := bluemonday.UGCPolicy().SanitizeBytes(output)

	t, err := template.New("mdp").Parse(defaultTemplate)
	if err != nil {
		return nil, err
	}
	if fname != "" {

		t, err = template.ParseFiles(fname)

		if err != nil {
			return nil, err
		}

	}

	c := content{

		Title: "MarkDown preview tool",
		Body:  template.HTML(body),
	}

	var buffer bytes.Buffer

	if err := t.Execute(&buffer, c); err != nil {
		return nil, err
	}

	// buffer.WriteString(header)
	// buffer.Write(body)
	// buffer.WriteString(footer)

	return buffer.Bytes(), nil
}

func saveHTML(outputfile string, htmlData []byte) error {
	// r w x
	// 4 2 1
	return os.WriteFile(outputfile, htmlData, 0644)
}
