// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {

	outline(os.Args[1:])

}

func outline(arg []string) error {
	resp, err := http.Get(arg[0])
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	//!+call
	result := ElementByID(doc, arg[1])
	fmt.Printf("result = %v", result)
	//!-call

	return nil
}

func ElementByID(doc *html.Node, id string) *html.Node {
	var ret *html.Node

	search := func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == "id" && a.Val == id {
					ret = n // ここで共有する.
					return false
				}
			}
		}
		return true
	}

	forEachNode(doc, search, nil)

	return ret
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) {
	if pre != nil {
		iscontinue := pre(n)
		if iscontinue == false {
			return
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		iscontinue := post(n)
		if iscontinue == false {
			return
		}
	}
}

//!-forEachNode
