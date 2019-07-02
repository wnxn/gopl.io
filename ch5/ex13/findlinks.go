// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 139.

// Findlinks3 crawls the web, starting with the URLs on the command line.
package main

import (
	"fmt"
	"github.com/adonovan/gopl.io/ch5/links"
	"log"
	"net/url"
	"os"
	"path"
	"strings"
)

//!+breadthFirst
// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

//!-breadthFirst

//!+crawl
func crawl(urlStr string) []string {
	normalUrl, err := url.Parse(urlStr)
	if err != nil{
		return []string{}
	}
	err = os.MkdirAll(path.Join("tmp", normalUrl.Host),os.ModePerm)
	if err != nil && os.IsNotExist(err){
		log.Println(err.Error())
		return []string{}
	}
	file, err :=os.Create(path.Join("tmp", normalUrl.Host, strings.ReplaceAll(strings.Trim(normalUrl.Path,"/"),"/",
		"-")))
	if err != nil{
		log.Println(err.Error())
	}
	defer file.Close()
	fmt.Printf("Host: %s, Path: %s\n", normalUrl.Host, normalUrl.Path)
	list, err := links.Extract(urlStr)
	file.WriteString(strings.Join(list, "\n"))
	if err != nil {
		log.Print(err)
	}
	return list
}

//!-crawl

//!+main
func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	//dirPath := path.Join(os.Getenv("GOPATH") ,"src/github.com/adonovan/gopl.io")
	breadthFirst(crawl, []string{"https://golang.org"})
}
//!-main
