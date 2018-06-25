package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	var address = "localhost:8080"
	if len(os.Args) == 3 {
		address = os.Args[1] + ":" + os.Args[2]
	}
	log.Printf("Connecting to %s\n", address)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		done<-struct{}{}
	}()
	mustCopy(conn, os.Stdin)
	conn.CloseWrite()
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
