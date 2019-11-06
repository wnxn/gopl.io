// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 254.
//!+

// Chat is a server that lets clients chat with each other.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

var timeout = 10 * time.Second

//!+broadcaster
type client chan<- string // an outgoing message channel

type clientSet struct {
	ch    client
	name  string
	conn  net.Conn
	timer *time.Timer
}

func (c *clientSet) Reset() {
	c.timer.Reset(timeout)
}

var (
	entering = make(chan clientSet)
	leaving  = make(chan clientSet)
	messages = make(chan string) // all incoming client messages
)

func disconnect(cli clientSet) {
	select {
	case <-cli.timer.C:
		cli.conn.Close()
	}
}

func broadcaster() {
	clients := make(map[clientSet]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli.ch <- msg
				cli.Reset()
			}

		case cli := <-entering:
			for curCli := range clients {
				cli.ch <- fmt.Sprintf("%s already in this chat", curCli.name)
			}
			clients[cli] = true
			go disconnect(cli)
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

//!-broadcaster

//!+handleConn
func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"

	newClient := clientSet{
		ch:    ch,
		name:  who,
		conn:  conn,
		timer: time.NewTimer(timeout),
	}

	entering <- newClient

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- newClient
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

//!+main
func main() {
	listener, err := net.ListenIP("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

//!-main
