// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 221.
//!+

// Netcat1 is a read-only TCP client.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
)

var (
	newyork = flag.String("NewYork", "", "")
	tokyo   = flag.String("Tokyo", "", "")
	london  = flag.String("London", "", "")
)

func main() {
	flag.Parse()
	go MyDial("NY", *newyork)
	go MyDial("TKY", *tokyo)
	MyDial("LON", *london)

}

func MyDial(city string, address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for timeS := ""; ; fmt.Fscanf(conn, "%s", &timeS) {
		fmt.Println(city, timeS)
	}
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

//!-
