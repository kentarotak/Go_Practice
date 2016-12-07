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
	"time"
)

//!+broadcaster
type client chan<- string // an outgoing message channel

var (
	entering   = make(chan client)
	leaving    = make(chan client)
	messages   = make(chan string) // all incoming client messages
	addnames   = make(chan string)
	rmnames    = make(chan string)
	rcvmessage = make(chan struct{})
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients

	var username []string
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			clients[cli] = true
			// 今存在するユーザーの表示.
			cli <- "\n現在以下のユーザがいます\n"
			for _, name := range username {
				cli <- name
			}

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		case name := <-addnames:
			username = append(username, name)
		case name := <-rmnames:
			remove(username, name)
		}
	}
}

//!-broadcaster

//!+handleConn
func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch
	// 名前を追加.
	addnames <- who

	go func() {
		ticker := time.NewTicker(300 * time.Second)
		flag := false
		for {
			select {
			case <-ticker.C:
				fmt.Printf("時間\n")
				if flag == false {
					fmt.Printf("close\n")
					conn.Close()
					ticker.Stop()
				}
				flag = false
			case <-rcvmessage:
				fmt.Printf("message受信\n")
				flag = true
			}
		}
	}()

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
		rcvmessage <- struct{}{}
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- ch
	messages <- who + " has left"
	//名前を削除.
	rmnames <- who
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
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func remove(numbers []string, search string) []string {
	result := []string{}
	for _, num := range numbers {
		if num != search {
			result = append(result, num)
		}
	}
	return result
}

//!-main
