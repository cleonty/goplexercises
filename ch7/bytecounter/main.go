package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	reader := bytes.NewReader(p)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*c++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return len(p), nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	reader := bytes.NewReader(p)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		*c++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return len(p), nil
}

type MyByteCounter struct {
	wr    io.Writer
	count int
}

func (c *MyByteCounter) Write(p []byte) (int, error) {
	count, err := c.wr.Write(p)
	if err != nil {
		return count, err
	}
	c.count += count
	return count, nil
}

func CountingWriter(w io.Writer) (io.Writer, *int) {
	counter := &MyByteCounter{w, 0}
	return counter, &counter.count
}

func main() {
	var bc ByteCounter
	var wc WordCounter
	var lc LineCounter

	bc.Write([]byte("hello"))
	wc.Write([]byte("hello"))
	lc.Write([]byte("hello"))

	fmt.Println(bc, wc, lc)

	bc, wc, lc = 0, 0, 0
	name := "Dolly"
	fmt.Fprintf(&bc, "hello, %s\nbye %s\n", name, name)
	fmt.Fprintf(&wc, "hello, %s\nbye %s\n", name, name)
	fmt.Fprintf(&lc, "hello, %s\nbye %s\n", name, name)
	fmt.Println(bc, wc, lc)

	writer, pCount := CountingWriter(os.Stdout)
	fmt.Fprintf(writer, "hello %s\n", name)
	fmt.Println(*pCount)
}
