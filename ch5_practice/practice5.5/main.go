package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	var words, images int
	for _, url := range os.Args[1:] {
		words, images, _ = CountWordsAndImages(url)
	}

	fmt.Printf("words=%d, images=%d", words, images)

}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML")
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	words, images = visit(n)
	return
}

func visit(n *html.Node) (words, images int) {

	if n.Type == html.TextNode || n.Type == html.DocumentNode {
		words = len(strings.Split(n.Data, " "))
	}

	if n.Data == "img" {
		images++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		tmpWords, tempImages := visit(c)
		words += tmpWords
		images += tempImages
	}

	return
}
