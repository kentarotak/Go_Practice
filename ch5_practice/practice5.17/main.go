// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args[1:]) >= 2 {
		resp, err := http.Get(os.Args[1])
		if err != nil {
			os.Exit(1)
		}
		defer resp.Body.Close()
		doc, err := html.Parse(resp.Body)
		if err != nil {
			os.Exit(1)
		}
		result := ElementByTagName(doc, "h1")

		for _, a := range result {
			fmt.Printf("-----------------\n")
			fmt.Printf("%#v\n", a)
			fmt.Printf("-----------------\n")
		}

	} else {
		fmt.Printf("Too few Argments\n")
	}
}

func ElementByTagName(doc *html.Node, ids ...string) []*html.Node {
	result := forEachNode(doc, nil, isElementByTagName, ids...)

	return result
}

func forEachNode(n *html.Node, pre, post func(n *html.Node, ids ...string) bool, ids ...string) []*html.Node {

	if pre != nil {
		pre(n, ids...)
	}

	var result []*html.Node

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result = append(result, forEachNode(c, pre, post, ids...)...)
	}

	var isMatchElement bool
	if post != nil {
		isMatchElement = post(n, ids...)
		if isMatchElement == true {
			result = append(result, n)
		}
	}

	return result
}

//!-forEachNode

func isElementByTagName(n *html.Node, ids ...string) bool {

	for _, id := range ids {
		//fmt.Printf("%s, %s\n", n.Data, id)
		if id == n.Data {
			return true
		}
	}

	return false
}
