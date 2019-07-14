package main

import (
	"bufio"
	"bytes"
	"fmt"
	"unicode"
)

func main() {
	var w WordCounter
	//fmt.Fprint(&w, "  hello   world  ass d ")
	w.Write2([]byte("  hello   world  ass d "))
	fmt.Println(w)

	var l LineCounter
	//fmt.Fprint(&l, "sefef\nedsfef")
	l.Write2([]byte("sefef\nedsfef"))
	fmt.Println(l)
}

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	inWord := false
	for _, v := range string(p) {
		if inWord == false && unicode.IsLetter(v) {
			inWord = true
		}
		if inWord == true && unicode.IsSpace(v) {
			inWord = false
			*c++
		}
	}
	if inWord == true {
		*c++
	}
	return int(*c), nil
}

func (c *WordCounter) Write2(p []byte) (int, error) {
	s := bufio.NewScanner(bytes.NewReader(p))
	s.Split(bufio.ScanWords)
	count := 0
	for s.Scan() {
		count++
	}
	*c = WordCounter(count)
	return count, nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	if len(p) != 0 {
		*c++
	}
	for _, v := range string(p) {
		if v == '\n' {
			*c++
		}
	}
	return int(*c), nil
}

func (c *LineCounter) Write2(p []byte) (int, error) {
	s := bufio.NewScanner(bytes.NewReader(p))
	s.Split(bufio.ScanLines)
	for s.Scan() {
		*c++
	}
	return int(*c), nil
}
