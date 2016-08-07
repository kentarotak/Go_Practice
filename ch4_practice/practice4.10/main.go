// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 112.
//!+

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

//!+
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	nowtime := time.Now()
	fmt.Printf("%d\n", nowtime.Year())

	var periodOver1year []*github.Issue
	var periodUnder1year []*github.Issue
	var periodUnder1month []*github.Issue
	for _, item := range result.Items {
		if nowtime.Year()-item.CreatedAt.Year() > 1 {
			//1年以上経過.
			periodOver1year = append(periodOver1year, item)
		} else {
			if nowtime.Month()-item.CreatedAt.Month() > 1 {
				// １年未満.
				periodUnder1year = append(periodUnder1year, item)
			} else {
				periodUnder1month = append(periodUnder1month, item)
			}
		}
	}

	fmt.Printf("★★１年以上経過のIssues★★\n")
	for _, item := range periodOver1year {
		// デバック用
		//fmt.Println(item.CreatedAt)

		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}

	fmt.Printf("★★１年未満かつ1か月以上経過のIssues★★\n")
	for _, item := range periodUnder1year {
		// デバック用
		//fmt.Println(item.CreatedAt)

		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}

	fmt.Printf("★★1か月未満のIssues★★\n")
	for _, item := range periodUnder1month {
		// デバック用
		//fmt.Println(item.CreatedAt)

		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}

}

//!-

/*
//!+textoutput
$ go build gopl.io/ch4/issues
$ ./issues repo:golang/go is:open json decoder
13 issues:
#5680    eaigner encoding/json: set key converter on en/decoder
#6050  gopherbot encoding/json: provide tokenizer
#8658  gopherbot encoding/json: use bufio
#8462  kortschak encoding/json: UnmarshalText confuses json.Unmarshal
#5901        rsc encoding/json: allow override type marshaling
#9812  klauspost encoding/json: string tag not symmetric
#7872  extempora encoding/json: Encoder internally buffers full output
#9650    cespare encoding/json: Decoding gives errPhase when unmarshalin
#6716  gopherbot encoding/json: include field name in unmarshal error me
#6901  lukescott encoding/json, encoding/xml: option to treat unknown fi
#6384    joeshaw encoding/json: encode precise floating point integers u
#6647    btracey x/tools/cmd/godoc: display type kind of each named type
#4237  gjemiller encoding/base64: URLEncoding padding is optional
//!-textoutput
*/
