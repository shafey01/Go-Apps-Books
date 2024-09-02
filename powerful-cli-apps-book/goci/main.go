// • Building the program using go build to verify if the program structure is
// valid.
// • Executing tests using go test to ensure the program does what it’s intended
// to do.
// • Executing gofmt to ensure the program’s format conforms to the standards.
// report erratum • discussChapter 6. Controlling Processes • 164
// • Executing git push to push the code to the remote shared Git repository
// that hosts the program code.
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	projDir := flag.String("p", "", "project directory")
	flag.Parse()

	if err := run(*projDir, os.Stdout); err != nil {
		os.Exit(1)
	}
}

func run(projDir string, out io.Writer) error {
	if projDir == "" {
		return fmt.Errorf("project directory is required: %w", ErrValidation)

	}
	args := []string{"build", ".", "errors"}

	cmd := exec.Command("go", args...)
	cmd.Dir = projDir

	if err := cmd.Run(); err != nil {
		return &stepErr{step: "go build", msg: "go build failed", cause: err}
	}

	_, err := fmt.Fprintln(out, "Go build: Success")
	return err

}
