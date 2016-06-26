// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo3 prints its command-line arguments.
package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

// テスト用変数.
var out io.Writer = os.Stdout

//!+
func main() {
	echo(os.Args[1:])
}

// 引数の数によって、コマンドラインにindexとその引数を改行して表示する.
func echo(args []string) error {

	var s, sep string
	for i := 0; i < len(args); i++ {
		s += sep + strconv.Itoa(i) + " " + args[i]
		sep = "\n"
	}

	fmt.Fprintf(out, s)
	fmt.Fprintln(out)

	return nil
}

//!-
