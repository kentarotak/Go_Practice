package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"

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
	//fmt.Printf("%#v\n", n)
	if n.Type == html.TextNode {
		if regexp.MustCompile(`.*\n|.*\t`).Match([]byte(n.Data)) == false {
			fmt.Println(n.Data)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}

	return links

}

func check_regexp(reg, str string) {
	fmt.Println(regexp.MustCompile(reg).Match([]byte(str)))
}
