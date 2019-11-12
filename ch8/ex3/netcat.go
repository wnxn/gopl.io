// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 227.

// Netcat is a simple read/write client for TCP servers.
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

//!+
func main() {
	conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{[]byte{127, 0, 0, 1}, 8000, ""})
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		num, err := io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		if err != nil{
			log.Println(num,err)
		}

		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)
	err = conn.CloseWrite()
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println("close write")
	<-done // wait for background goroutine to finish
	conn.CloseRead()

}

//!-

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
