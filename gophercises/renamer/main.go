package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

var (
	filePattern = `^(christmas 2016) \((\d+) of (\d+)\)\.(\w+)$`
)

func main() {
	dirPath := flag.String("dir", ".", "Directory to rename files into it")
	flag.Parse()

	if err := filepath.Walk(*dirPath, walkFn); err != nil {
		log.Fatalf("can't walk into the dire %d", err)
	}
}

func walkFn(path string, info fs.FileInfo, err error) error {

	if info.IsDir() {
		return nil
	}
	regmatches := regexp.MustCompile(filePattern)
	namesMatched := regmatches.FindStringSubmatch(info.Name())
	if len(namesMatched) == 0 {
		return nil
	}
	var (
		epsiodename   = namesMatched[1]
		epsiodenumber = namesMatched[2]
		epsiodeexten  = namesMatched[4]

		renamed = fmt.Sprintf("Epsiode %s: %s.%s", epsiodenumber, epsiodename, epsiodeexten)
	)
	return os.Rename(path, filepath.Join(filepath.Dir(path), renamed))

}
