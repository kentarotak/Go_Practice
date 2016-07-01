package main

import (
	"bytes"
	"testing"
)

// 数値のテスト.
func TestNumEcho(t *testing.T) {

	testfile := []string{"1", "9", "8", "100", "11"}
	want := "0 1\n1 9\n2 8\n3 100\n4 11\n"
	out = new(bytes.Buffer)
	result := out.(*bytes.Buffer)

	echo(testfile)

	if result.String() != want {
		t.Errorf("NumTestError Output: %s", result.String())
	}
}

// 文字列のテスト.
func TestStrEcho(t *testing.T) {

	testfile := []string{"apple", "orange", "pine", "banana", "lemon"}
	want := "0 apple\n1 orange\n2 pine\n3 banana\n4 lemon\n"
	out = new(bytes.Buffer)
	result := out.(*bytes.Buffer)

	echo(testfile)

	if result.String() != want {
		t.Errorf("StrTestError Output: %s", result.String())
	}
}
