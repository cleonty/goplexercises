package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
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

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "Вы " + who
	messages <- who + " подключился"
	entering <- ch
	input := bufio.NewScanner(conn)

	time.After(time.Second * 10)
	alive := make(chan struct{})
	go checkClient(who, conn, alive)
	for input.Scan() {
		alive <- struct{}{}
		messages <- who + ": " + input.Text()
	}
	leaving <- ch
	messages <- who + " отключился"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func checkClient(who string, conn net.Conn, alive <-chan struct{}) {
	for {
		select {
		case <-time.After(time.Second * 10):
			messages <- who + " будет отключен по таймауту"
			conn.Close()
			return
		case <-alive:
			messages <- who + " еще здесь"
		}
	}
}
