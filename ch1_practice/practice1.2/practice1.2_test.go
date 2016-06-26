package main

import (
	"bytes"
	"testing"
)

func TestEcho(t *testing.T) {

	testfile := []string{"1", "9", "8", "100", "11"}
	want := "0 1\n1 9\n2 8\n3 100\n4 11 \n"
	out = new(bytes.Buffer)
	result := out.(*bytes.Buffer)

	echo(testfile)

	if result.String() != want {
		t.Errorf("input: 10 100 100\\n, Output: %s", result.String())
	}

}
