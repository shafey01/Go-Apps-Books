package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {

	lines := flag.Bool("l", false, "Count Lines")
	flag.Parse()

	fmt.Println(count(os.Stdin, *lines))

}

func count(r io.Reader, linesFlag bool) int {

	scanner := bufio.NewScanner(r)

	if !linesFlag {

		scanner.Split(bufio.ScanWords)
	}
	wc := 0

	for scanner.Scan() {
		wc++
	}

	return wc

}
