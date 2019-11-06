// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 222.

// Clock is a TCP server that periodically writes the time.
package main

import (
	"flag"
	"fmt"
	"io"
	"k8s.io/klog"
	"log"
	"net"
	"os"
	"time"
)

func handleConn(c net.Conn, location string) {
	defer c.Close()
	_, err := time.LoadLocation(location)
	if err != nil {
		klog.Error(err.Error())
		return
	}
	for {
		_, err = io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			klog.Error(err.Error())
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

var (
	port = flag.Int("port", 0, "port of clock server")
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()
	klog.Info(*port)
	timezone := os.Getenv("TZ")
	klog.Infof("timezone=%s", timezone)
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatal(err)
	}
	//!+
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn, timezone) // handle connections concurrently
	}
	//!-
}
