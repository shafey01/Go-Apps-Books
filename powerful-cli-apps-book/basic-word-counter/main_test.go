package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {

	bytes := bytes.NewBufferString("use when\n scanning and the \n maximum size\n")
	expected := 3
	result := count(bytes, true)

	if result != expected {

		t.Errorf("Expected %d, got %d instead.\n", expected, result)

	}

}
