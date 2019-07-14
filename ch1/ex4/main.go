// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	filenames := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "", counts, filenames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, arg, counts, filenames)
			f.Close()
		}
	}
	for line, n := range counts {
		if n >= 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, strings.Join(filenames[line], ","))
		}
	}
}

func countLines(f *os.File, filename string, counts map[string]int, filenames map[string][]string) {
	input := bufio.NewScanner(f)
	dupCurrentFile := make(map[string]int)
	for input.Scan() {
		counts[input.Text()]++
		dupCurrentFile[input.Text()]++
		if dupCurrentFile[input.Text()] == 1 {
			filenames[input.Text()] = append(filenames[input.Text()], filename)
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}

//wangxin
//!-
