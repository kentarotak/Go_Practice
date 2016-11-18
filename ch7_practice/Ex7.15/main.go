// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 195.

// Http4 is an e-commerce server that registers the /list and /price
// endpoint by calling http.HandleFunc.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/kentarotak/Go_Practice/ch7_practice/Ex7.15/eval"
)

//!+main

func main() {

	fmt.Printf("式を入力してください\n")

	in := bufio.NewScanner(os.Stdin)

	var s string
	if in.Scan() {
		s = in.Text()
	}

	expr, err := eval.Parse(s)

	if err != nil {
		fmt.Printf("構文が誤っています %s \n", err)
	}

	vars := make(map[eval.Var]float64)

	eval.GetLiteral(vars, expr)

	for key, _ := range vars {
		fmt.Printf("変数 %s の値を入れてください\n", key)
		if in.Scan() {
			num, err := strconv.ParseFloat(in.Text(), 64)
			if err != nil {
				fmt.Printf("数値を入れてください")
				os.Exit(1)
			}
			vars[key] = num
		}
	}

	fmt.Printf("答えは %f です\n", expr.Eval(vars))

}
