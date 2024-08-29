package main

import (
	"io"
	"path/filepath"
	"testing"
)

func BenchmarkRun(b *testing.B) {
	filenames, err := filepath.Glob("./testdata/*.csv")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := run(filenames, "avg", 2, io.Discard); err != nil {
			b.Error(err)
		}
	}

}
