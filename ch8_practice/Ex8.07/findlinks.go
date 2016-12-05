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
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"gopl.io/ch5/links"
)

func crawl(url_data string) []string {
	url_data = strings.Trim(url_data, "/")
	u, _ := url.Parse(url_data)
	//fmt.Printf("ROW = %s\n", url_data)
	//fmt.Printf("Host = %s\n", u.Host)
	//fmt.Printf("Path = %s\n", u.Path)

	// ファイルのDL
	response, err := http.Get(url_data)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, filename := path.Split(url_data)

	// dirの作成.
	p := strings.Split(u.Path, "/")

	var dir string
	if len(p) > 2 {
		for i := 0; i < len(p)-1; i++ {
			dir += p[i] + "/"
		}
		dir = strings.Trim(dir, "/")

		if len(p) == 3 {
			os.Mkdir(dir, 0777)
		} else {
			os.MkdirAll(dir, 0777)
		}
	} else if len(p) == 2 {
		dir = p[1]
		os.Mkdir(dir, 0777)
	}

	file, err := os.OpenFile(dir+"/"+filename, os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		file.Close()
	}()

	file.Write(body)

	// ここまでDL

	// Link先を解析
	list, err := links.Extract(url_data)
	if err != nil {
		log.Print(err)
	}

	//listから他ドメインの要素を除去.
	var result []string
	for _, val := range list {
		tmp, _ := url.Parse(val)
		if tmp.Host == u.Host {
			result = append(result, val)
		}
	}

	return result
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
