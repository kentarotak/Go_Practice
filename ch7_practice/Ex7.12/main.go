// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 195.

// Http4 is an e-commerce server that registers the /list and /price
// endpoint by calling http.HandleFunc.
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

//!+main
var elementsList = template.Must(template.New("dataElement").Parse(`
<table border="1">
<tr style='text-align: left'>
	<th>Item</th>
	<th>Price</th>
</tr>
{{range .}}
<tr>
	<td>{{.Item}}
	<td>{{.Price}}
</tr>
{{end}}
</table>
`))

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!-main

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

type dataElement struct {
	Item  string
	Price float32
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	var elem []dataElement
	for item, price := range db {
		var temp dataElement
		temp.Item = item
		temp.Price = float32(price)
		elem = append(elem, temp)
	}

	if err := elementsList.Execute(w, elem); err != nil {
		log.Fatal(err)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	if item != "" && price != "" {
		val, err := strconv.ParseFloat(price, 32)
		if err != nil {
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "can't create: invalid price\n")
			return
		}
		db[item] = dollars(val)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "can't create: meybe typos\n")
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if _, ok := db[item]; ok {
		val, err := strconv.ParseFloat(price, 32)
		if err != nil {
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "can't create: invalid price\n")
			return
		}

		db[item] = dollars(val)
		fmt.Fprintf(w, "update end\n")
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "can't delete: not exsist element\n")
		return
	}
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; ok {
		delete(db, item)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "can't delete: not exsist element\n")
		return
	}
}
