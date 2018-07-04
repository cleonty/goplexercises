// Exercise 8.8
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	var address = "localhost:8080"
	if len(os.Args) > 1 {
		address = os.Args[1]
	}
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		log.Printf("client connected")
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	var wg sync.WaitGroup
	done := false;
	lines := make(chan string)
	go func() {
		for input.Scan() {
			lines <- input.Text()
		}
	}()
	for ; !done; {
		select {
		case line := <-lines:
			wg.Add(1)
			go func() {
				defer wg.Done()
				echo(c, line, 1*time.Second)
			}()
		case <-time.After(5 * time.Second):
			done = true;
		}
	}
	wg.Wait()
	bye(c);
	c.Close()
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func bye(c net.Conn) {
	fmt.Fprintln(c, "\t", "bye")
}

