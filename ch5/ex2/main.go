// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 122.
//!+main

// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func fetch(url string)io.Reader{
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	return resp.Body
}

func main() {
	doc, err := html.Parse(fetch("https://golang.org"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	num :=  populate(doc)
	fmt.Println(num)
}

//!-main

//!+visit
// visit appends to links each link found in n and returns the result.
func populate( n *html.Node) int {
	res := 0
	if n == nil{
		return res
	}
	if n.Type == html.ElementNode && (n.Data == "p" || n.Data == "div" || n.Data == "span") {
		res += len(n.Attr)
		fmt.Println(n.Data, n.Attr)
	}
	return populate(n.FirstChild)+populate(n.NextSibling) + res
}

//!-visit

/*
//!+html
package html

type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}

type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}

func Parse(r io.Reader) (*Node, error)
//!-html
*/
