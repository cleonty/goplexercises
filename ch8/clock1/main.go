package main

import (
	"io"
	"log"
	"net"
	"os"
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
		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:15.999\n"))
		if err != nil {
			log.Print(err)
			return
		}
		time.Sleep(10 * time.Millisecond)
	}
}
