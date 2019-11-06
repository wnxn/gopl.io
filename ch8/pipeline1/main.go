// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 228.

// Pipeline1 demonstrates an infinite 3-stage pipeline.
package main

import "fmt"

//!+
func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func(ch chan<- int) {
		for x := 0; x < 10; x++ {
			ch <- x
		}
		close(ch)
	}(naturals)

	// Squarer
	go func(ch1 chan<- int, ch2 <-chan int) {
		for x := range ch2 {
			ch1 <- x * x
		}
		close(ch1)
	}(squares, naturals)

	// Printer (in main goroutine)

	for x := range squares {
		fmt.Println(x)
	}
}

//!-
