package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

var trackList = template.Must(template.New("trackList").Parse(`
<table>
<tr style='text-align: left'>
    <th><a href='http://localhost:8000/?=title'>Title</a></th>
    <th><a href='http://localhost:8000/?=artist'>Artist</a></th>
    <th><a href='http://localhost:8000/?=album'>Album</a></th>
    <th><a href='http://localhost:8000/?=year'>Year</a></th>
    <th><a href='http://localhost:8000/?=length'>Length</a></th>
</tr>
{{range .}}
<tr>
    <td>{{.Title}}</td>
    <td>{{.Artist}}</td>
    <td>{{.Album}}</td>
    <td>{{.Year}}</td>
    <td>{{.Length}}</td>
</tr>    
{{end}}
</table>
`))

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(os.Stdout, "URL_query= %q\n", r.URL.RawQuery)

		if r.URL.RawQuery != "" {
			sortDataTable(r.URL.RawQuery)
			printTracksHTML(w)
		} else {
			printTracksHTML(w)
		}
	}
	http.HandleFunc("/", handler)
	//!-http
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return
}

func printTracksHTML(out io.Writer) {
	if err := trackList.Execute(out, tracks); err != nil {
		log.Fatal(err)
	}
}

func sortDataTable(query string) {

	switch {
	case query == "=title":
		StableSortByMyself(byTitle(tracks))
	case query == "=year":
		StableSortByMyself(byYear(tracks))
	case query == "=artist":
		StableSortByMyself(byArtist(tracks))
	case query == "=length":
		StableSortByMyself(byLength(tracks))
	case query == "=album":
		StableSortByMyself(byAlbum(tracks))
	}
}
