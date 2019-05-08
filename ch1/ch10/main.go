// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 17.
//!+

// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	filename := "result"
	fopen, err := os.OpenFile(filename,os.O_RDWR|os.O_CREATE | os.O_APPEND, 0755)
	if err != nil{
		if os.IsExist(err){
			fopen, _ = os.Create(filename)
		}
		fmt.Print(err)
	}
	if err != nil{
		fmt.Print(err)
	}
	for range os.Args[1:] {
		fmt.Fprintf(fopen,"%s\n",<-ch) // receive from channel ch
	}
	fmt.Fprintf(fopen,"%.2fs elapsed\n", time.Since(start).Seconds())
	fopen.Close()
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

//!-
