// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 224.

// Reverb2 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {

	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

//!+
func handleConn(c net.Conn, ch chan<- int) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
		ch <- len(input.Text())
	}
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}

//!-

func main() {
	l, err := net.ListenTCP("tcp4", &net.TCPAddr{[]byte{0, 0, 0, 0}, 8000, ""})
	if err != nil {
		log.Fatal(err)
	}
	ch := make(chan int)
	go func(ch2 <-chan int) {
		for {
			fmt.Printf("main: %d bytes\n", <-ch2)
		}
	}(ch)

	for {
		conn, err := l.AcceptTCP()
		var wg sync.WaitGroup
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go func() {
			wg.Add(1)
			defer wg.Done()
			handleConn(conn, ch)
		}()
		go func() {
			wg.Wait()
			fmt.Println("close write")
			conn.CloseWrite()
		}()
	}

}
