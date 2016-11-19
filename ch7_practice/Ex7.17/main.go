// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 214.
//!+

// go run main.go p.copyright < w3c.txt

// result
/*
p.copyright 以下の要素が抽出できている
html body div p a : Copyright
html body div p :  © 2006
html body div p a acronym : W3C
html body div p sup : ®
html body div p :  (
html body div p a acronym : MIT
html body div p : ,
html body div p a acronym : ERCIM
html body div p : ,
html body div p a : Keio
html body div p : ), All Rights Reserved. W3C
html body div p a : liability
html body div p : ,
html body div p a : trademark
html body div p :  and
html body div p a : document use
html body div p :  rules apply.
*/

// Xmlselect prints the text of selected elements of an XML document.
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []xml.StartElement // stack of element names
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				var str string
				for _, elem := range stack {
					str += elem.Name.Local + " "
				}
				fmt.Printf("%s: %s\n", str, tok)
			}
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x []xml.StartElement, y []string) bool {

	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}

		// ここを変える
		// y[0]を構文ルールに従ってパースする. #:id, .:class
		parsedata, tag := parseInput(y[0])
		//fmt.Printf("parsedata = %v, tag = %s\n", parsedata, tag)
		if tag == "none" {
			if x[0].Name.Local == parsedata[0] {
				y = y[1:]
			}
		} else {
			if x[0].Name.Local == parsedata[0] {
				for _, attr := range x[0].Attr {
					//fmt.Printf("attr %s\n", attr.Name.Local)
					//fmt.Printf("attr %s\n", attr.Value)
					if attr.Name.Local == tag && attr.Value == parsedata[1] {
						y = y[1:]
					}
				}
			}
		}

		//
		x = x[1:]
	}
	return false
}

func parseInput(val string) (result []string, tag string) {

	var parsedata []string

	tag = "none"
	if strings.Contains(val, ".") {
		parsedata = strings.Split(val, ".")
		tag = "class"
	} else if strings.Contains(val, "#") {
		parsedata = strings.Split(val, "#")
		tag = "id"
	} else {
		parsedata = append(parsedata, val)
	}

	return parsedata, tag

}

//!-
