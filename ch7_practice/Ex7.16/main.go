// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 195.

// Http4 is an e-commerce server that registers the /list and /price
// endpoint by calling http.HandleFunc.
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/kentarotak/Go_Practice/ch7_practice/Ex7.15/eval"
)

type Data map[eval.Var]float64

var expr eval.Expr

//!+main

type ShowVariable struct {
	Function string
	Variable []string
}

func (d Data) inputVariable(w http.ResponseWriter, req *http.Request) {
	Item := req.URL.RawQuery

	tmp := strings.Split(Item, "=")

	if len(tmp) < 1 {
		log.Fatal("入力が誤っています")
	}

	expr, _ = eval.Parse(tmp[1])
	//fmt.Fprintf(w, "%v\n", expr)

	eval.GetLiteral(d, expr)

	//fmt.Fprintf(w, "%v\n", d)
	var Val ShowVariable

	Val.Function = tmp[1]

	for tmp := range d {
		Val.Variable = append(Val.Variable, string(tmp))
	}
	sort.Strings(Val.Variable)

	variabletmpl, _ := template.ParseFiles("variable.tmpl")
	if err := variabletmpl.Execute(w, Val); err != nil {
		log.Fatal(err)
	}

}

func (d Data) input(w http.ResponseWriter, req *http.Request) {
	maintmpl, _ := template.ParseFiles("index.tmpl")
	if err := maintmpl.ExecuteTemplate(w, "index", nil); err != nil {
		log.Fatal(err)
	}
}

func (d Data) showResult(w http.ResponseWriter, req *http.Request) {
	Item := req.URL.Query().Get("item")
	tmp := strings.Split(Item, ",")

	tmp = tmp[1:]

	fmt.Fprintf(w, "%v :: %d\n", tmp, len(tmp))

	var key []string
	for tmp := range d {
		key = append(key, string(tmp))
	}

	sort.Strings(key)
	fmt.Fprintf(w, "%v\n", key)

	for i := 0; i < len(key); i++ {
		num, err := strconv.ParseFloat(tmp[i], 64)
		if err != nil {
			fmt.Fprintf(w, "入力が誤っています\n")
			return
		}
		d[eval.Var(key[i])] = num
	}

	fmt.Fprintf(w, "答えは %f です\n", expr.Eval(eval.Env(d)))

}

func main() {
	var data Data
	data = make(map[eval.Var]float64)
	http.HandleFunc("/input", data.input)
	http.HandleFunc("/variable", data.inputVariable)
	http.HandleFunc("/result", data.showResult)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
