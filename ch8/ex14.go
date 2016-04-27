// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 254.
//!+

// Chat is a server that lets clients chat with each other.
// exercise 8-14
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type client struct {
	msg chan<- string // an outgoing message channel
	who string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

// cactiveClients returns active client names
func activeClients(cs map[client]bool) []string {
	whos := make([]string, 0, len(cs))
	for k, b := range cs {
		if b {
			whos = append(whos, k.who)
		}
	}
	return whos
}

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli.msg <- msg
			}

		case cli := <-entering:
			clients[cli] = true
			cli.msg <- fmt.Sprintf("current clients: %v", activeClients(clients))

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.msg)
		}
	}
}

//!-broadcaster

//!+handleConn
func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	input := bufio.NewScanner(conn)

	var who string
	fmt.Fprint(conn, "What's your name?: ")
	if input.Scan() {
		who = input.Text()
	}

	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- client{ch, who}

	in := make(chan struct{})
	// close non active connection
	go func() {
		ticker := time.NewTicker(5 * time.Minute)
		for {
			select {
			case <-ticker.C:
				conn.Close()
			case <-in:
				// message coming! reset timer.
				ticker.Stop()
				ticker = time.NewTicker(5 * time.Minute)
			}
		}
	}()
	for input.Scan() {
		messages <- who + ": " + input.Text()
		in <- struct{}{}
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- client{ch, who}
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
	listener, err := net.Listen("tcp", "localhost:8000")
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
