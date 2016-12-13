// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 254.
//!+

// Chat is a server that lets clients chat with each other.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

//!+broadcaster
type client chan<- string // an outgoing message channel

type cmdSet struct {
	cmd string
	ch  chan string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan cmdSet) // all incoming client messages

)

func cmdHandler() {
	clients := make(map[client]bool) // all connected clients

	for {
		select {
		case cli := <-entering:
			cli <- "220 FTP server Ready\n"
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		case msg := <-messages:
			interPretationCmd(msg.ch, msg.cmd)
		}
	}
}

func interPretationCmd(ch chan<- string, cmd string) {

	data := strings.Split(cmd, " ")

	//fmt.Printf("%s %s\n", data[0], data[1])

	switch data[0] {
	case "USER":
		ch <- "331 User name okay, need password"
	case "PASS":
		ch <- "230 User logged in, proceed"
	case "PORT":
		// 相手先のIPアドレスとポート番号を用いて、コネクションを張る.
		//conn, err := net.Dial("tcp", "localhost:8000")
		ch <- "200 PORT command successful"
	}

}

//!+handleConn
func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		cmd := input.Text()
		fmt.Printf("コマンド: %s\n", cmd)

		var val cmdSet

		val.cmd = cmd
		val.ch = ch
		messages <- val
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- ch
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

//!+main
func main() {
	listener, err := net.Listen("tcp", "localhost:21")
	if err != nil {
		log.Fatal(err)
	}

	go cmdHandler()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
