// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 222.

// Clock is a TCP server that periodically writes the time.
package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		str := time.Now().Format("15:04:05\n")
		str = strings.Trim(str, "\n")
		fmt.Fprintln(c, str)
		//_, err := io.WriteString(c, str)
		/*
			if err != nil {
				return // e.g., client disconnected
			}
		*/
		time.Sleep(1 * time.Second)
	}
}

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("引数が足りません\n")
		os.Exit(1)
	}
	port := os.Args[1]

	host := "localhost:" + port

	listener, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatal(err)
	}
	//!+
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}
	//!-
}
