package main

import (
	"fmt"
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
	defer conn.Close()
	var s string
	for {
		if _, err := fmt.Fscanf(conn, "%s", &s); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", s)
	}

	//mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {

	}
}
