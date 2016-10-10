// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 139.

// Findlinks3 crawls the web, starting with the URLs on the command line.
package main

import (
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

//!+breadthFirst
// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

//!-breadthFirst

//!+crawl
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

//!-crawl

//!+main
func main() {

	breadthFirst(crawl, os.Args[1:])
}

//!-main
