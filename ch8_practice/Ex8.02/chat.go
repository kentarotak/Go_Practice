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
	"os"
	"os/exec"
	"strconv"
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

var clients = make(map[client]bool) // all connected clients
var clientsConnection = make(map[client]net.Conn)
var rootpath = make(map[client]string)

func cmdHandler() {

	for {
		select {
		case cli := <-entering:
			clients[cli] = true
			rootpath[cli], _ = os.Getwd() //接続時のディレクトリを仮想rootpathにする.
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

	fmt.Printf("%s\n", data[0])

	switch data[0] {
	case "USER":
		string := fmt.Sprintf("331 Password required for %s", data[1])
		ch <- string
	case "PASS":
		ch <- "230 User logged in, proceed"
	case "PORT":
		// 相手先のIPアドレスとポート番号を用いて、コネクションを張る.
		fmt.Printf("%s\n", data[1])
		prs := strings.Split(data[1], ",")

		stringIp := fmt.Sprintf("%s.%s.%s.%s", prs[0], prs[1], prs[2], prs[3])

		highbyte, _ := strconv.Atoi(prs[4])
		lowbyte, _ := strconv.Atoi(prs[5])

		portnum := highbyte*256 + lowbyte

		addr := fmt.Sprintf("%s:%d", stringIp, portnum)
		fmt.Printf("%s\n", addr)
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			log.Fatal(err)
		}

		clientsConnection[ch] = conn

		ch <- "200 PORT command successful"
	case "NLST":
		ch <- "150 File status okay; about to open data connection."

		// 外部コマンドを使用する.
		out, _ := exec.Command("ls").CombinedOutput()

		// 別コネクションを用いて、フォルダ情報を送信.
		fmt.Fprintln(clientsConnection[ch], string(out))

		ch <- "226 Closing data connection."
		clientsConnection[ch].Close()
	case "SYST":
		ch <- "215 Windows\n"
	case "FEAT":
		ch <- "211 "
	case "XPWD":
		pwd, _ := os.Getwd()
		pwd = strings.TrimPrefix(pwd, rootpath[ch])
		string := fmt.Sprintf("257 \"/%s\" is current directory.", pwd)
		ch <- string
	case "TYPE A":
		ch <- "200 TYPE SET TO A"
	case "CWD":
		ch <- "250 CWD command successful."
	default:
		ch <- "504 Command not implemented for that parameter."
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
