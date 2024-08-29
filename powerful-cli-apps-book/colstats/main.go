package main

import (
	"flag"
	"fmt"
	"io"
	"os"
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
	for _, filename := range filenames {
		f, err := os.Open(filename)
		if err != nil {
			return fmt.Errorf("can not open file %w", err)
		}
		data, err := csvtofloat(f, column)
		if err != nil {
			return err
		}

		if err := f.Close(); err != nil {
			return err
		}

		consolidate = append(consolidate, data...)
	}
	_, err := fmt.Fprintln(out, opFunc(consolidate))
	return err
}
