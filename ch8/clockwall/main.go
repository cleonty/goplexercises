package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

// Conn represents a tcp connection
type Conn struct {
	address string
	conn    net.Conn
	err     error
}

func (conn *Conn) printTime() {
	if conn.err == nil {
		var s string
		_, err := fmt.Fscanf(conn.conn, "%s", &s)
		if err != nil {
			log.Println(err)
			conn.err = err
		}
		log.Printf("%s: %s\n", conn.address, s)
	}
}

func main() {
	var conns []Conn
	for _, addr := range os.Args[1:] {
		c, err := net.Dial("tcp", addr)
		if err != nil {
			log.Println(err)
			continue
		}
		conn := Conn{addr, c, nil}
		defer c.Close()
		conns = append(conns, conn)
	}
	for {
		for _, c := range conns {
			c.printTime()
		}
	}
}
