// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 221.
//!+

// Netcat1 is a read-only TCP client.
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	var buf []*net.Conn
	var area []string
	for _, data := range os.Args[1:] {
		area_tmp, port := parseArg(data)
		var tmp *net.Conn
		go connect(port, tmp)
		area = append(area, area_tmp)
		buf = append(buf, tmp)

	}

	fmt.Fprintf(os.Stdout, "%s \t %s\n", area[0], area[1])
	for {
		time.Sleep(2 * time.Second)
		io.Copy(os.Stdout, *(buf[0]))
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

func connect(port string, buf *net.Conn) {
	conn, err := net.Dial("tcp", port)

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	buf = &conn
	//mustCopy(write, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

//!-
