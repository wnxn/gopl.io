// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 243.

// Crawl3 crawls web links starting with the command-line arguments.
//
// This version uses bounded parallelism.
// For simplicity, it does not address the termination problem.
//
package main

import (
	"flag"
	"fmt"
	"gopl.io/ch5/links"
	"log"
)

var (
	depth = flag.Int("depth", 1, "")
	link  = flag.String("link", "https://www.sina.com/", "")
)

type InternetLink struct {
	link  string
	depth int
}

func init() {
	flag.Parse()
}

func crawl(url InternetLink) (res []InternetLink) {
	fmt.Println(url)
	list, err := links.Extract(url.link)
	if err != nil {
		log.Print(err)
	}

	for _, v := range list {
		res = append(res, InternetLink{v, url.depth + 1})
	}
	return res
}

//!+
func main() {
	fmt.Println(*depth)
	worklist := make(chan []InternetLink)  // lists of URLs, may have duplicates
	unseenLinks := make(chan InternetLink) // de-duplicated URLs

	// Add command-line arguments to worklist.
	go func() { worklist <- []InternetLink{{*link, 0}} }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 40; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}
	i := 0
	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[InternetLink]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] && link.depth <= *depth {
				seen[link] = true
				i++
				fmt.Println(i, link)
				unseenLinks <- link
			}
		}
	}
}

//!-
