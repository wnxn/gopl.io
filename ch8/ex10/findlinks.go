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
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

var done = make(chan struct{})

func cancel() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func crawl(url string) []string {
	if cancel() {
		panic("wd")
	}

	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

//!+
func main() {
	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs

	// Add command-line arguments to worklist.
	go func() { worklist <- os.Args[1:] }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				if cancel() {
					return
				}
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}
	i := 0
	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		close(done)
	}()
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if cancel() {
				panic("hi")
			}
			if !seen[link] {
				seen[link] = true
				i++
				fmt.Println(i, link)
				unseenLinks <- link
			}
		}
	}
}

//!-
