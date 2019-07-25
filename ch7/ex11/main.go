// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 195.

// Http4 is an e-commerce server that registers the /list and /price
// endpoint by calling http.HandleFunc.
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"html/template"
)

//!+main

var databaseList = template.Must(template.New("databaseList").Parse(`
<h1>database</h1>
<table>
<tr style='text-align: left'>
  <th>Item</th>
  <th>Price</th>
</tr>
{{range .}}
<tr>
  <td>{{.Item}}</a></td>
  <td>{{.Price}}</a></td>
</tr>
{{end}}
</table>
`))

func main() {
	db := database{
		entity:map[string]dollars{"shoes": 50, "socks": 5},
		}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/read",db.read)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!-main

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database struct{
	entity map[string]dollars
	mut sync.Mutex
}

type table struct{
	Item string
	Price dollars
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	t := make([]table,0)
	for item, price := range db.entity {
		t = append(t, table{item,price})
	}
	databaseList.Execute(w,t)
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db.entity[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

// uri: /update?item=socks&price=60
func (db database) update(w http.ResponseWriter, req *http.Request){
	item := req.URL.Query().Get("item")
	if _,ok:=db.entity[item];ok{
		price, err := strconv.Atoi(req.URL.Query().Get("price"))
		if err != nil{
			http.Error(w, err.Error(), http.StatusNotAcceptable)
		}
		db.mut.Lock()
		defer db.mut.Unlock()
		db.entity[item] = dollars(price)
		fmt.Fprintf(w, "succeed to add %s %d", item, price)
	}else{
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

// uri: /delete?item=socks
func (db database)delete(w http.ResponseWriter, req *http.Request){
	item := req.URL.Query().Get("item")
	if _, ok := db.entity[item]; ok {
		db.mut.Lock()
		defer db.mut.Unlock()
		delete(db.entity, item)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

// uri: /create?item=socks&price=20
func (db database) create(w http.ResponseWriter, req *http.Request){
	item := req.URL.Query().Get("item")
	if _, ok:=db.entity[item]; ok{
		http.Error(w, "entity already existed", http.StatusCreated)
	}else{
		price, err := strconv.Atoi(req.URL.Query().Get("price"))
		if err != nil{
			http.Error(w, err.Error(), http.StatusNotAcceptable)
		}
		db.mut.Lock()
		defer db.mut.Unlock()
		db.entity[item] = dollars(price)
		fmt.Fprintf(w, "succeed to add %s %d", item, price)
	}
}

// uri: /read?item=socks
func (db database) read(w http.ResponseWriter, req *http.Request){
	item := req.URL.Query().Get("item")
	if price, ok := db.entity[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

