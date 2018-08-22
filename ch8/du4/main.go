package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// walkDir рекурсивно обходит дерево файлов с корнем в dir
// и отправляет размер каждого найденного файла в fileSizes.
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	if cancelled() {
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}

}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}:
	case <-done:
		return nil
	}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

var verbose = flag.Bool("v", false, "вывод промежуточных результатов")

func main() {
	flag.Parse()
	roots := flag.Args()
	fmt.Println(flag.Args())
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	go func() {
		os.Stdin.Read(make([]byte, 1))
		fmt.Println("close done channel")
		close(done)
	}()
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()
	var nfiles, nbytes int64
loop:
	for {
		select {
		case <-done:
			fmt.Println("cancelled, wait for the channel is empty")
			for range fileSizes {

			}
			fmt.Println("the channel is empty, done")
			// trick: see if there is only one goroutine
			// panic("end of program")
			return
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d файлов %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
