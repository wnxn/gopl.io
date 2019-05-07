// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 16.
//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	const http_prefix = "http://"
	for _, url := range os.Args[1:] {
		if ! strings.HasPrefix(url, http_prefix){
			url = http_prefix + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		//b, err := ioutil.ReadAll(resp.Body)
		// ex1.7
		io.Copy(os.Stdout, resp.Body)
		// ex1.8
		status := resp.Status
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v, status code %s\n", url, err, status)
			os.Exit(1)
		}
		// ex1.9
		fmt.Fprintf(os.Stdout, "status=%s\n",status)
		//fmt.Printf("%s", b)
	}
}

//!-
