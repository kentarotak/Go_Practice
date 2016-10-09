// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	//!+call
	forEachNode(doc, startElement, endElement)
	//!-call

	return nil
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {

	// 空要素の場合は、省略形を用いる.
	if isEmptyElement(n) == true {
		fmt.Printf("%*s<%s", (depth+1)*2, "", n.Data)

		for _, a := range n.Attr {
			fmt.Printf(" %s=\"%s\"", a.Key, a.Val)
		}
		fmt.Printf("/>\n")
	} else {
		if pre != nil {
			pre(n)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			forEachNode(c, pre, post)
		}

		if post != nil {
			post(n)
		}
	}
}

//!-forEachNode

//!+startend
var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s", depth*2, "", n.Data)
		for _, a := range n.Attr {
			fmt.Printf(" %s=\"%s\"", a.Key, a.Val)
		}
		fmt.Printf(">\n")
		depth++
	} else if n.Type == html.TextNode {
		if regexp.MustCompile(`.*[A-Za-z_\[\]]`).Match([]byte(n.Data)) == true {
			words := strings.Split(n.Data, "\n")
			for _, word := range words {
				fmt.Printf("%*s%s\n", depth*2, "", word)
			}
		}
	} else if n.Type == html.DocumentNode {
		if regexp.MustCompile(`.*[A-Za-z_\[\]]`).Match([]byte(n.Data)) == true {
			fmt.Printf("%*s%s\n", depth*2, "", n.Data)
		}
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

//!-startend

func isEmptyElement(n *html.Node) bool {

	emptyelements := []string{"area", "base", "basefont", "br", "col", "frame", "hr", "img",
		"input", "isindex", "link", "meta", "param"}

	for _, element := range emptyelements {
		if element == n.Data {
			return true
		}
	}

	return false
}
