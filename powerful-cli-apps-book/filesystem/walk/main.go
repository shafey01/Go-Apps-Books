package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

// • -root: The root of the directory tree to start the search.
//    The default is the current directory.
// • -list: List files found by the tool.
//   When specified, no other actions will be executed.
// • -ext: File extension to search.
//    When specified, the tool will only match files
//    with this extension.
// • -size: Minimum file size in bytes.
//    When specified, the tool will only match
//   files whose size is larger than this value.
//

// config struct

type config struct {
	list    bool
	ext     string
	size    int64
	del     bool
	wLog    io.Writer
	archive string
}

func main() {

	// flags
	root := flag.String("root", ".", "Root path to start")
	list := flag.Bool("list", false, "Just list files")
	ext := flag.String("ext", "", "File extension to search")
	size := flag.Int64("size", 0, "File size to search")
	del := flag.Bool("del", false, "Delete file")
	logFile := flag.String("log", "", "Log delted files")
	archive := flag.String("archive", "", "Archive directory")
	flag.Parse()

	// Open file for writing logs
	var (
		f   = os.Stdout
		err error
	)
	// Open file for reading
	if *logFile != "" {
		f, err = os.OpenFile(*logFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer f.Close()
	}

	// instant from config struct
	c := config{
		list:    *list,
		ext:     *ext,
		size:    *size,
		del:     *del,
		wLog:    f,
		archive: *archive,
	}

	// Run function

	if err := Run(*root, os.Stdout, c); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Run(root string, out io.Writer, cfg config) error {

	delLogger := log.New(cfg.wLog, "DELETED File: ", log.LstdFlags)
	return filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if filterOut(path, cfg.ext, cfg.size, info) {
				return nil
			}

			if cfg.list {
				return listFile(path, out)
			}

			if cfg.archive != "" {
				if err := archiveFile(cfg.archive, root, path); err != nil {
					return err
				}
			}
			if cfg.del {
				return delFile(path, delLogger)
			}
			return listFile(path, out)
		})
}
