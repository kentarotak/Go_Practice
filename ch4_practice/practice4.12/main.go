package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type XkcdIndex struct {
	Month      string      `json:"month"`
	Num        json.Number `json:"num"`
	Year       json.Number `json:"year"`
	SafeTitle  string      `json:"safe_title"`
	Transcript string      `json:"transcript"`
	Alt        string      `json:"alt"`
	Img        string      `json:"img"`
	Title      string      `json:"title"`
	Day        json.Number `json:"day"`
}

// JSON Interface.
//http://xkcd.com/info.0.json 最新のコミック
//http://xkcd.com/614/info.0.json (comic #614)

func main() {
	cmd := os.Args[1:]
	// indexロード.
	objs := loadIndex()
	// 最新のjsonと比較.
	var first, last int
	if objs == nil {
		first = 0
	} else {
		first, _ = strconv.Atoi(string(objs[len(objs)-1].Num))
	}
	last, _ = chklatestindex()

	if first == last {
		fmt.Printf("indexは最新です!\n")
	} else {
		fmt.Printf("取得のindexは%dで最新は%dなので、未取得のindexを取得します\n", first, last)

		for i := first + 1; i <= last; i++ {
			//未取得indexの取得.
			result, err := downLoadIndex(i)
			if err != nil {
				fmt.Printf("err %s", err)
				continue
			}
			objs = append(objs, result)
		}

		// listの更新
		saveIndex(objs)
	}

	// 検索されたワードをもとにURLと内容を表示する.
	searchInfo(cmd, objs)
}

func searchInfo(cmd []string, objs []XkcdIndex) {
	for _, obj := range objs {
		var result int
		// SafeTitle, Transcript, Titleを検索し、いずれかの文字列に
		//検索用語が存在した場合は標準出力にURLとTranscriptを表示する
		result = strings.Index(obj.SafeTitle, cmd[0])
		if result != -1 {
			showInfo(obj)
			continue
		}
		result = strings.Index(obj.Transcript, cmd[0])
		if result != -1 {
			showInfo(obj)
			continue
		}
		result = strings.Index(obj.Title, cmd[0])
		if result != -1 {
			showInfo(obj)
			continue
		}

	}
}

func showInfo(obj XkcdIndex) {
	fmt.Printf("------検索にヒットした要素-------------\n")
	fmt.Printf("URL: https://xkcd.com/%s/\n", obj.Num)
	fmt.Printf("Transcript: %s\n", obj.Transcript)
	fmt.Printf("-----------------------------------\n")
}

func saveIndex(objs []XkcdIndex) {
	file, _ := os.OpenFile("index.txt", os.O_WRONLY, 0644)
	fmt.Fprintf(file, "[")

	size := len(objs)
	for i, obj := range objs {
		bdata, _ := json.Marshal(obj)
		if i != size-1 {
			fmt.Fprintf(file, "%s,\n", bdata)
		} else {
			fmt.Fprintf(file, "%s\n", bdata)
		}
	}
	fmt.Fprintf(file, "]")

	file.Close()
}

func loadIndex() []XkcdIndex {
	file, err := ioutil.ReadFile("index.txt")
	if err != nil {
		fmt.Printf("you must create index.txt %d\n", err)
		os.Exit(1)
	}

	var result []XkcdIndex
	json_err := json.Unmarshal(file, &result)
	if err != nil {
		fmt.Println("Format Error: ", json_err)
	}
	return result
}

func chklatestindex() (int, error) {

	resp, err := http.Get("http://xkcd.com/info.0.json")

	if err != nil {
		fmt.Printf("Can't get data: %s\n", err)
		os.Exit(1)
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("search query failed: %s\n", resp.Status)
		resp.Body.Close()
		os.Exit(1)
	}

	var latestjson XkcdIndex
	if err := json.NewDecoder(resp.Body).Decode(&latestjson); err != nil {
		resp.Body.Close()
		fmt.Printf("Can't Decode: %s\n", err)
		os.Exit(1)
	}

	resp.Body.Close()
	index, _ := strconv.Atoi(string(latestjson.Num))
	return index, nil

}

func downLoadIndex(indexNum int) (XkcdIndex, error) {

	var index XkcdIndex
	url := "http://xkcd.com/" + strconv.Itoa(indexNum) + "/info.0.json"

	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("Can't get data: %s\n", err)
		os.Exit(1)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return index, fmt.Errorf("search query failed: %s\n", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&index); err != nil {
		resp.Body.Close()
		fmt.Printf("Can't Decode: %s\n", err)
		os.Exit(1)
	}

	resp.Body.Close()

	return index, nil
}
