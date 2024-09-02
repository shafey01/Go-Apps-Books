package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

// • encoding/csv: To read data from CSV files.
// • fmt: To print formatted results out.
// • io: To provide the io.Reader interface.
// • strconv: To convert string data into numeric data.

func sum(data []float64) float64 {
	sum := 0.0

	for _, v := range data {
		sum += v
	}

	return sum
}

func avg(data []float64) float64 {
	return sum(data) / float64(len(data))
}

type statsFunc func(data []float64) float64

func csvtofloat(r io.Reader, columns int) ([]float64, error) {

	reader := csv.NewReader(r)
	reader.ReuseRecord = true

	columns--

	// allData, err := reader.Read()
	// if err != nil {
	// 	return nil, fmt.Errorf("can not read file from file: %w", err)

	// }

	var data []float64
	for i := 0; ; i++ {
		// Read the file record by record
		row, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("can't read the file %w", err)
		}

		if i == 0 {
			continue
		}

		if len(row) <= columns {
			return nil, fmt.Errorf("%w: File has only %d columns", ErrInvalidColumn, len(row))

		}

		v, err := strconv.ParseFloat(row[columns], 64)
		if err != nil {
			return nil, fmt.Errorf("%w: %s", ErrNotNumber, err)
		}

		data = append(data, v)
	}
	return data, nil

}
