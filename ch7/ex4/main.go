// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 122.
//!+main

// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"os"
	"strings"
)

type mystring struct {
	s string
	i int64 // current reading index
}

func (r *mystring) Read(b []byte) (n int, err error) {
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	n = copy(b, r.s[r.i:])
	r.i += int64(n)
	return
}

func NewReader(s string) io.Reader {
	return &mystring{
		s: content,
		i: 0,
	}
}

func main() {
	doc, err := html.Parse(strings.NewReader(content))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

//!-main

//!+visit
// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
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

var content string = `<!DOCTYPE html>
<html lang="en">
<meta charset="utf-8">
<meta name="description" content="Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.">
<meta name="viewport" content="width=device-width, initial-scale=1">
<meta name="theme-color" content="#00ADD8">

  <title>The Go Programming Language</title>

<link type="text/css" rel="stylesheet" href="/lib/godoc/style.css">

<link rel="search" type="application/opensearchdescription+xml" title="godoc" href="/opensearch.xml" />


<script>window.initFuncs = [];</script>

<script>
var _gaq = _gaq || [];
_gaq.push(["_setAccount", "UA-11222381-2"]);
window.trackPageview = function() {
  _gaq.push(["_trackPageview", location.pathname+location.hash]);
};
window.trackPageview();
window.trackEvent = function(category, action, opt_label, opt_value, opt_noninteraction) {
  _gaq.push(["_trackEvent", category, action, opt_label, opt_value, opt_noninteraction]);
};
</script>

<script src="/lib/godoc/jquery.js" defer></script>



<script src="/lib/godoc/playground.js" defer></script>

<script>var goVersion = "go1.12.7";</script>
<script src="/lib/godoc/godocs.js" defer></script>

<div id="lowframe" style="position: fixed; bottom: 0; left: 0; height: 0; width: 100%; border-top: thin solid grey; background-color: white; overflow: auto;">
...
</div><!-- #lowframe -->

<body class="Site">
<header class="Header js-header">
  <nav class="Header-nav ">
    <a href="/"><img class="Header-logo" src="/lib/godoc/images/go-logo-blue.svg" alt="Go"></a>
    <button class="Header-menuButton js-headerMenuButton" aria-label="Main menu" aria-expanded="false">
      <div class="Header-menuButtonInner">
    </button>
    <ul class="Header-menu">
      <li class="Header-menuItem"><a href="/doc/">Documents</a></li>
      <li class="Header-menuItem"><a href="/pkg/">Packages</a></li>
      <li class="Header-menuItem"><a href="/project/">The Project</a></li>
      <li class="Header-menuItem"><a href="/help/">Help</a></li>
      
      <li class="Header-menuItem Header-menuItem--search">
        <form class="HeaderSearch" role="search" action="/search">
          <input class="HeaderSearch-input"
                type="search"
                name="q"
                placeholder="Search"
                aria-label="Search"
                autocapitalize="off"
                autocomplete="off"
                autocorrect="off"
                spellcheck="false"
                required>
          <button class="HeaderSearch-submit">
            <!-- magnifying glass: --><svg class="HeaderSearch-icon" width="24" height="24" viewBox="0 0 24 24"><title>Search</title><path d="M15.5 14h-.79l-.28-.27C15.41 12.59 16 11.11 16 9.5 16 5.91 13.09 3 9.5 3S3 5.91 3 9.5 5.91 16 9.5 16c1.61 0 3.09-.59 4.23-1.57l.27.28v.79l5 4.99L20.49 19l-4.99-5zm-6 0C7.01 14 5 11.99 5 9.5S7.01 5 9.5 5 14 7.01 14 9.5 11.99 14 9.5 14z"/><path d="M0 0h24v24H0z" fill="none"/></svg>
          </button>
        </form>
      </li>
    </ul>
  </nav>
</header>

<main id="page" class="Site-content">
<div class="container">
`
