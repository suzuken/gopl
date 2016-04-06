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

func handleTCPConn(c *net.TCPConn) {
	var wg sync.WaitGroup
	input := bufio.NewScanner(c)
	for input.Scan() {
		go func() {
			defer wg.Done()
			wg.Add(1)
			echo(c, input.Text(), 1*time.Second)
		}()
	}
	go func() {
		wg.Wait()
		c.Close()
	}()
}

func handleConn(c net.Conn) {
	tc, ok := c.(*net.TCPConn)
	if !ok {
		log.Fatal("connection should be tcp connection. failed.")
	}
	handleTCPConn(tc)
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
