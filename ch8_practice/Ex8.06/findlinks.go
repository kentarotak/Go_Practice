// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 243.

// Crawl3 crawls web links starting with the command-line arguments.
//
// This version uses bounded parallelism.
// For simplicity, it does not address the termination problem.
//
package main

import (
	"flag"
	"fmt"
	"log"

	"gopl.io/ch5/links"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

type Buff struct {
	url string
	cnt int
}

//!+
func main() {
	worklist := make(chan []Buff)  // lists of URLs, may have duplicates
	unseenLinks := make(chan Buff) // de-duplicated URLs

	var f = flag.Int("depth", 0, "url depth")
	flag.Parse()

	// 深さ
	depth := *f

	// Add command-line arguments to worklist.
	var buff []Buff
	for _, elm := range flag.Args() {
		var tmp Buff
		tmp.url = elm
		tmp.cnt = 0
		buff = append(buff, tmp)
	}

	go func() { worklist <- buff }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				if link.cnt < depth {
					foundLinks := crawl(link.url)
					var buff []Buff
					for _, elm := range foundLinks {
						var tmp Buff
						tmp.url = elm
						tmp.cnt = link.cnt + 1 // 深さカウンタをインクリメント.
						buff = append(buff, tmp)
					}
					go func() { worklist <- buff }()
				}
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)

	for list := range worklist {
		for _, link := range list {
			if !seen[link.url] {
				seen[link.url] = true
				unseenLinks <- link
				fmt.Printf("url = %s : depth = %d\n", link.url, link.cnt)
			}
		}
	}

}

//!-
