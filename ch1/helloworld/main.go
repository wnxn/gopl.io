// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 1.

// Helloworld is our first Go program.
//!+
package main

import "fmt"

var test map[int]string = map[int]string{
	4:"he",
}

func main() {
	fmt.Println("Hello, 世界")
}

//!-
