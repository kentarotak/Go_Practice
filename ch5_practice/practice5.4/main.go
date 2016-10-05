package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {

	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		doc, err := html.Parse(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		for _, link := range visit(nil, doc) {
			fmt.Println(link)
		}
	}

}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && (n.Data == "a" || n.Data == "img" || n.Data == "script") {
		for _, a := range n.Attr {
			if a.Key == "href" {
				tmp := n.Data + ":" + a.Val
				links = append(links, tmp)
			} else if a.Key == "src" {
				tmp := n.Data + ":" + a.Val
				links = append(links, tmp)
			}
		}
	} else if n.Type == html.ElementNode && n.Data == "link" {
		isStyleSheet := false
		for _, a := range n.Attr {
			// relがstylesheetのときに、hrefを取得する.
			if a.Key == "rel" && a.Val == "stylesheet" {
				isStyleSheet = true
			} else if a.Key == "href" {
				if isStyleSheet == true {
					tmp := "StyleSheet" + ":" + a.Val
					links = append(links, tmp)
				}
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
