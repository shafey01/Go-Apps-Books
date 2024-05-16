package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {

	bytes := bytes.NewBufferString("use when scanning and the maximum size\n")
	expected := 7
	result := count(bytes)

	if result != expected {

		t.Errorf("Expected %d, got %d instead.\n", expected, result)

	}

}
