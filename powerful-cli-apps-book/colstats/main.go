package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
)

func main() {
	op := flag.String("op", "sum", "Opertion to be executed")
	column := flag.Int("col", 1, "CSV column on which to be executed opertion")
	flag.Parse()

	if err := run(flag.Args(), *op, *column, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// 1. filenames of type []string: A slice of strings representing the file names to
// process.
// 2. op of type string: A string representing the operation to execute, such as
// sum or average.
// 3. column of type int: An integer representing the column on which to execute
// the operation.
// 4. out of type io.Writer: An io.Writer interface to print out the results. By using
// the interface, you can print to STDOUT in the program while allowing tests
// to capture results using a buffer.

func run(filenames []string, op string, column int, out io.Writer) error {

	var opFunc statsFunc

	if len(filenames) == 0 {
		return ErrNoFiles
	}

	if column < 1 {
		return fmt.Errorf("%w: %d", ErrInvalidColumn, column)
	}

	switch op {
	case "sum":
		opFunc = sum
	case "avg":
		opFunc = avg
	default:
		return fmt.Errorf("%w: %s", ErrInvalidOperation, op)
	}

	consolidate := make([]float64, 0)
	resCh := make(chan []float64)
	errCh := make(chan error)
	doneCh := make(chan struct{})
	fileCh := make(chan string)
	wg := sync.WaitGroup{}

	go func() {
		defer close(fileCh)
		for _, fname := range filenames {
			fileCh <- fname
		}
	}()

	for i := 0; i < runtime.NumCPU(); i++ {
		fmt.Println(runtime.NumCPU())
		wg.Add(1)
		go func() {

			defer wg.Done()
			for fname := range fileCh {
				f, err := os.Open(fname)
				if err != nil {

					errCh <- fmt.Errorf("can not open file %w", err)
					return
				}
				data, err := csvtofloat(f, column)
				if err != nil {
					errCh <- err
					return
				}

				if err := f.Close(); err != nil {
					errCh <- err
					return
				}
				resCh <- data
			}
		}()

	}
	go func() {
		wg.Wait()
		close(doneCh)
	}()

	for {
		select {
		case err := <-errCh:
			return err
		case data := <-resCh:
			consolidate = append(consolidate, data...)
		case <-doneCh:
			_, err := fmt.Fprintln(out, opFunc(consolidate))

			return err
		}
	}
}
