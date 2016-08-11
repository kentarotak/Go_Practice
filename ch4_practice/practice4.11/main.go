package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"gopl.io/ch4/github"
)

// コマンドルール.
// 第一引数は -○○ で実行条件を指定.
// 第二引数以降は ○○:要素　で条件を指定する.

// Issue作成
// -create title:test body:this is test repos:kentarotak/Go_Practice token:○○
// Issue更新
// -edit title:test body:this is test repos:kentarotak/Go_Practice number:8 token:○○
// Issueクローズ
// -edit title:test body:this is test repos:kentarotak/Go_Practice number:8 status:close token:○○
// 読みだし
// -search repo:kentarotak/Go_Practice

//!+
func main() {
	// 入力部.
	cmd := os.Args[1:]
	if cmd[0] == "-execEditer" {
		cmd = execEditer(cmd[1:])
	}

	// searchはそのまま使う.
	if cmd[0] == "-search" {
		result, err := github.SearchIssues(cmd[1:])

		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d issues:\n", result.TotalCount)
		for _, item := range result.Items {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	} else {
		// create(作成), update(更新), close(クローズ)は入力をパースする.
		parsedata := make(map[string]string)
		parseCmd(cmd[1:], parsedata)

		// 作成、更新、クローズの実施.
		if cmd[0] == "-create" {
			createIssue(parsedata)
		}

		if cmd[0] == "-edit" {
			editIssue(parsedata)
		}

		if cmd[0] == "-lock" {
			lockIssue(parsedata)
		}
	}

}

func parseCmd(cmd []string, parsedata map[string]string) {
	var keys []int
	for i, data := range cmd {
		if strings.Contains(data, ":") == true {
			keys = append(keys, i)
		}
	}

	for i := 0; i < len(keys); i++ {
		var diff int
		if i != len(keys)-1 {
			diff = keys[i+1] - keys[i]
		} else {
			diff = 0
		}
		parse := strings.Split(cmd[keys[i]], ":")
		str := parse[1]
		for j := 1; j < diff; j++ {
			str += " " + cmd[keys[i]+j]
		}
		parsedata[parse[0]] = str
	}
}

func execEditer(cmd []string) []string {
	exe := exec.Command(cmd[0], "text.txt")
	exe.Stdin = os.Stdin
	exe.Stdout = os.Stdout
	exe.Run()

	f, err := os.Open("text.txt")

	if err != nil {
		fmt.Printf("err %d\n", err)
	}

	input := bufio.NewScanner(f)
	input.Split(bufio.ScanWords)

	var editcmd []string
	for input.Scan() {
		editcmd = append(editcmd, input.Text())
	}
	f.Close()

	if err := os.Remove("text.txt"); err != nil {
		fmt.Println(err)
	}
	return editcmd
}

type CreateIssue struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func createIssue(parsedata map[string]string) {

	_, ok := parsedata["title"]
	if ok != true {
		fmt.Printf("titleを設定してください\n")
		return
	}
	_, ok = parsedata["body"]
	if ok != true {
		fmt.Printf("bodyを設定してください\n")
		return
	}
	_, ok = parsedata["repos"]
	if ok != true {
		fmt.Printf("reposを設定してください\n")
		return
	}
	_, ok = parsedata["token"]
	if ok != true {
		fmt.Printf("tokenを設定してください\n")
		return
	}

	var issue = CreateIssue{Title: parsedata["title"], Body: parsedata["body"]}

	data, err := json.Marshal(issue)
	if err != nil {
		log.Fatalf("JSON marshaling Failed %s\n", err)
	}

	url := "https://api.github.com/repos/"
	url += parsedata["repos"] + "/issues"

	req, _ := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer(data),
	)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "token "+parsedata["token"])

	client := &http.Client{Timeout: time.Duration(15 * time.Second)}
	resp, err := client.Do(req)

	fmt.Printf("status = %d\n", resp.Status)
	defer resp.Body.Close()

}

type EditIssue struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	State string `json:"state"`
}

func editIssue(parsedata map[string]string) {

	_, ok := parsedata["title"]
	if ok != true {
		fmt.Printf("titleを設定してください\n")
		return
	}
	_, ok = parsedata["body"]
	if ok != true {
		fmt.Printf("bodyを設定してください\n")
		return
	}
	_, ok = parsedata["repos"]
	if ok != true {
		fmt.Printf("reposを設定してください\n")
		return
	}

	_, ok = parsedata["state"]
	if ok != true {
		// 指定がない場合はopen
		parsedata["state"] = "open"
	}

	_, ok = parsedata["token"]
	if ok != true {
		fmt.Printf("tokenを設定してください\n")
		return
	}
	_, ok = parsedata["number"]
	if ok != true {
		fmt.Printf("numberを設定してください\n")
		return
	}

	var issue = EditIssue{Title: parsedata["title"], Body: parsedata["body"], State: parsedata["state"]}

	data, err := json.Marshal(issue)
	if err != nil {
		log.Fatalf("JSON marshaling Failed %s\n", err)
	}

	url := "https://api.github.com/repos/"
	url += parsedata["repos"] + "/issues/" + parsedata["number"]

	req, _ := http.NewRequest(
		"PATCH",
		url,
		bytes.NewBuffer(data),
	)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "token "+parsedata["token"])

	client := &http.Client{Timeout: time.Duration(15 * time.Second)}
	resp, err := client.Do(req)

	fmt.Printf("status = %d\n", resp.Status)
	defer resp.Body.Close()

}

func lockIssue(parsedata map[string]string) {

	_, ok := parsedata["repos"]
	if ok != true {
		fmt.Printf("reposを設定してください\n")
		return
	}
	_, ok = parsedata["token"]
	if ok != true {
		fmt.Printf("tokenを設定してください\n")
		return
	}
	_, ok = parsedata["number"]
	if ok != true {
		fmt.Printf("numberを設定してください\n")
		return
	}

	url := "https://api.github.com/repos/"
	url += parsedata["repos"] + "/issues/" + parsedata["number"] + "/lock"

	req, _ := http.NewRequest(
		"PUT",
		url,
		nil,
	)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "token "+parsedata["token"])

	client := &http.Client{Timeout: time.Duration(15 * time.Second)}
	resp, _ := client.Do(req)

	fmt.Printf("status = %d\n", resp.Status)
	defer resp.Body.Close()

}
