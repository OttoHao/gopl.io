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
func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	textCh := make(chan string)

	go func() {
		for input.Scan() {
			textCh <- input.Text()
		}
	}()

	for {
		timeout := time.NewTimer(10 * time.Second)
		select {
		case text := <-textCh:
			go echo(c, text, 1*time.Second)
		case <-timeout.C:
			timeout.Stop()
			close(c)
			return
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}

func close(c net.Conn) {
	fmt.Println("close connection")
	if con, ok := c.(*net.TCPConn); ok {
		con.CloseWrite()
	} else {
		c.Close()
	}
}

//!-

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
