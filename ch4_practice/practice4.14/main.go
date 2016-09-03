package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/kentarotak/Go_Practice/ch4_practice/practice4.14/githubIssues"
)

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
　<th>BugDetail</th>
  <th>Milestone</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
  <td>{{.Body}}'</td>
  <td>{{.Milestone}}'</td>
</tr>
{{end}}
</table>
`))

func main() {

	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(os.Stdout, "URL_query= %q\n", r.URL.RawQuery)
		// ?q以降を取得する.
		query := strings.Split(r.URL.RawQuery, "/?q")
		if query[0] != "" {
			showHTML(w, query[0])
		} else {
			showMain(w)
		}
	}
	http.HandleFunc("/", handler)
	//!-http
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return

}

func showMain(out io.Writer) {
	maintmpl, _ := template.ParseFiles("index.tmpl")
	if err := maintmpl.ExecuteTemplate(out, "index", nil); err != nil {
		log.Fatal(err)
	}
}

func showHTML(out io.Writer, url string) {
	result, _ := githubIssues.SearchIssues(url)

	if result == nil {
		errormpl, _ := template.ParseFiles("error.tmpl")
		if err := errormpl.ExecuteTemplate(out, "error", nil); err != nil {
			log.Fatal(err)
		}
		return
	}

	if err := issueList.Execute(out, result); err != nil {
		log.Fatal(err)
	}
}
