package main

import (
	"fmt"
	"os"
	"testing"
)

// 空のtextが指定されたときに、
// mapの要素に何も入らずに終了すること
func TestStringEmpty(t *testing.T) {

	counts := make(map[string]int)
	filenames := make(map[string]string)
	f, err := os.Open("empty.txt")

	if err != nil {
		fmt.Fprintf(os.Stderr, "text not found", err)
	}

	filecountLines(f, counts, "empty.txt", filenames)

	if len(counts) != 0 {
		t.Errorf("TestStringEmpty error ")
	}

}

type testdata struct {
	key   string
	count int
}

// 第一引数に空のtext、第二引数に要素が入ったテキスト
// 第二引数に指定されたtextの要素が仕様に沿って出てくること
func TestStringEmptyAndElement(t *testing.T) {

	var tests = []struct {
		testval         string
		expect          int
		expect_filename string
	}{
		{"apple", 2, "\ttestfile.txt"},
		{"orange", 1, "\ttestfile.txt"},
	}

	counts := make(map[string]int)
	filenames := make(map[string]string)
	files := []string{"empty.txt", "testfile.txt"}

	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "text not found", err)
		}

		filecountLines(f, counts, arg, filenames)
	}

	for _, test := range tests {
		if got := counts[test.testval]; got != test.expect {
			t.Errorf("err %s, %v = %v", test.testval, test.expect, got)
		}
		if got := filenames[test.testval]; got != test.expect_filename {
			t.Errorf("file err %s, %v = %v", test.testval, test.expect_filename, got)
		}

	}

}

// 正常系 : testfile.txt, testfile2.txt を入力したときに
// 仕様に従って,mapファイルにデータが入力されること
func TestStringNormalTest(t *testing.T) {

	var tests = []struct {
		testval         string
		expect          int
		expect_filename string
	}{
		{"apple", 3, "\ttestfile.txt\ttestfile2.txt"},
		{"orange", 3, "\ttestfile.txt\ttestfile2.txt"},
		{"pine", 4, "\ttestfile2.txt"},
	}

	counts := make(map[string]int)
	filenames := make(map[string]string)
	files := []string{"testfile.txt", "testfile2.txt"}

	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "text not found", err)
		}

		filecountLines(f, counts, arg, filenames)
	}

	for _, test := range tests {
		if got := counts[test.testval]; got != test.expect {
			t.Errorf("err %s, %v = %v", test.testval, test.expect, got)
		}
		if got := filenames[test.testval]; got != test.expect_filename {
			t.Errorf("file err %s, %v = %v", test.testval, test.expect_filename, got)
		}
	}

}

// 正常系 ; testfile2.txt testfile.txtの順に引数を指定したときに
// 仕様に従って,mapファイルにデータが入力されること
func TestStringNormalTest2(t *testing.T) {

	var tests = []struct {
		testval         string
		expect          int
		expect_filename string
	}{
		{"apple", 3, "\ttestfile2.txt\ttestfile.txt"},
		{"orange", 3, "\ttestfile2.txt\ttestfile.txt"},
		{"pine", 4, "\ttestfile2.txt"},
	}

	counts := make(map[string]int)
	filenames := make(map[string]string)
	files := []string{"testfile2.txt", "testfile.txt"}

	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "text not found", err)
		}

		filecountLines(f, counts, arg, filenames)
	}

	for _, test := range tests {
		if got := counts[test.testval]; got != test.expect {
			t.Errorf("err %s, %v = %v", test.testval, test.expect, got)
		}
		if got := filenames[test.testval]; got != test.expect_filename {
			t.Errorf("file err %s, %v = %v", test.testval, test.expect_filename, got)
		}
	}

}
