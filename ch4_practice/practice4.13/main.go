package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type OmdMovieInfo struct {
	Title    string
	Response string
	Poster   string
}

// JSON Interface.
//http://www.omdbapi.com/?t=never+ending+story&y=&plot=short&r=json

func main() {
	cmd := os.Args[1:]

	var jsonobj OmdMovieInfo
	err := getJSON(cmd, &jsonobj)

	if err != nil {
		fmt.Printf("err: %s\n", err)
		os.Exit(1)
	}

	getPoster(&jsonobj)
}

func getPoster(jsonobj *OmdMovieInfo) {
	url := jsonobj.Poster

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("err: %s\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	file, err := os.Create(jsonobj.Title + ".jpg")
	if err != nil {
		fmt.Printf("err: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	io.Copy(file, resp.Body)

}

func getJSON(cmd []string, jsonobj *OmdMovieInfo) error {

	if len(cmd) == 0 {
		return errors.New("You Must type title")
	}
	var title string
	for i, data := range cmd {
		if i == 0 {
			title = data
		} else {
			title += "+" + data
		}
	}

	url := "http://www.omdbapi.com/?t=" + title + "&y=&plot=short&r=json"

	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
	}

	if err := json.NewDecoder(resp.Body).Decode(&jsonobj); err != nil {
		resp.Body.Close()
		return err
	}

	isSuccess := jsonobj.Response
	posterURL := jsonobj.Poster

	if isSuccess == "False" {
		return errors.New("Search nof found")
	}

	if posterURL == "not avaliable" {
		return errors.New("not avaliable")
	}

	resp.Body.Close()

	return nil
}
