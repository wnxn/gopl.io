// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 173.

// Bytecounter demonstrates an implementation of io.Writer that counts bytes.
package main

import (
	"fmt"
	"bufio"
	"flag"
	"os"
	"github.com/golang/glog"
	"io/ioutil"
	"bytes"
)

//!+bytecounter
func init(){
	flag.Set("logtostderr", "true")
	flag.Set("v","5")
}

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

//!-bytecounter
type WordCounter int

func (c *WordCounter) Write(p []byte)(int, error){
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan(){
		count++
	}
	*c = WordCounter(count)
	return count,scanner.Err()
}

type LineCounter int

func (c *LineCounter)Write(p []byte)(int, error){
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)
	count :=0
	for scanner.Scan(){
		count++
	}
	*c = LineCounter(count)
	return count,scanner.Err()
}

func main() {
	flag.Parse()
	//!+main
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // "5", = len("hello")

	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, Dolly")
	//!-main

	var w WordCounter
	w.Write([]byte("The order in which the methods appear is immaterial."))
	fmt.Println(w)

	w = 0
	var sentence = "The order in which the methods appear is immaterial."
	fmt.Fprintf(&w, "hello, %s", sentence)
	fmt.Println(w)

	filepath := os.Getenv("GOPATH")+"/src/github.com/adonovan/gopl.io/poem.txt"
	p, err := ioutil.ReadFile(filepath)
	if err != nil{
		glog.Fatal(err)
	}
	var l LineCounter
	l.Write(p)
	fmt.Println(l)

	l = 0
	fmt.Fprint(&l, "%s", string(p))
	fmt.Println(l)
}
