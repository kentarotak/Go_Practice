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
	if len(os.Args[1:]) <= 2 {
		resp, err := http.Get(os.Args[1])
		if err != nil {
			os.Exit(1)
		}
		defer resp.Body.Close()
		doc, err := html.Parse(resp.Body)
		if err != nil {
			os.Exit(1)
		}
		result := ElementById(doc, os.Args[2])

		fmt.Printf("%#v\n", result)
	} else {
		fmt.Printf("Too few Argments\n")
	}
}

func ElementById(doc *html.Node, id string) *html.Node {
	result := forEachNode(doc, id, isElementById, nil)

	return result
}

var isSearchEnd = true

func forEachNode(n *html.Node, id string, pre, post func(n *html.Node, id string) bool) *html.Node {

	if pre != nil {
		isSearchEnd = pre(n, id)
		if isSearchEnd == true {
			return n
		}
	}
	var result *html.Node
	for c := n.FirstChild; c != nil && isSearchEnd == false; c = c.NextSibling {
		result = forEachNode(c, id, pre, post)
	}

	if post != nil {
		post(n, id)
	}

	return result
}

//!-forEachNode

func isElementById(n *html.Node, id string) bool {

	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				return true
			}
		}
	}
	return false
}
