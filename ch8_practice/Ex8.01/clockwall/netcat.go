// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 221.
//!+

// Netcat1 is a read-only TCP client.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

var timedata = make(map[string]string)

func main() {

	for _, data := range os.Args[1:] {
		area, port := parseArg(data)
		go connect(area, port)
	}

	for {
		time.Sleep(2 * time.Second)
		var showstr string
		for area, time := range timedata {
			showstr += fmt.Sprintf("::Area %s: Time %s ::", area, time)
		}

		fmt.Printf("\r %s", showstr)

	}

}

func parseArg(in string) (area string, port string) {

	tmp := strings.Split(in, "=")

	if len(tmp) < 2 {
		log.Fatal("err")
	}
	area = tmp[0]
	port = tmp[1]

	return area, port
}

func connect(area string, port string) {
	conn, err := net.Dial("tcp", port)
	timedata[area] = ""

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Printf("接続!\n")

	input := bufio.NewScanner(conn)
	for input.Scan() {
		cmd := input.Text()
		timedata[area] = cmd
	}
}

//!-
