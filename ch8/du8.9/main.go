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
func walkDir(root int, dir string, n *sync.WaitGroup, sizeReponses chan<- SizeReponse) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(root, subdir, n, sizeReponses)
		} else {
			sizeReponses <- SizeReponse{root, entry.Size()}
		}
	}

}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

var verbose = flag.Bool("v", false, "вывод промежуточных результатов")

type SizeReponse struct {
	root   int
	nbytes int64
}

type SizeCount struct {
	nbytes int64
	nfiles int64
}

func main() {
	flag.Parse()
	roots := flag.Args()
	fmt.Println(flag.Args())
	if len(roots) == 0 {
		roots = []string{"."}
	}
	sizeResponses := make(chan SizeReponse)
	var n sync.WaitGroup
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	for i, root := range roots {
		n.Add(1)
		go walkDir(i, root, &n, sizeResponses)
	}
	go func() {
		n.Wait()
		close(sizeResponses)
	}()
	sizes := make([]SizeCount, len(roots))
loop:
	for {
		select {
		case resp, ok := <-sizeResponses:
			if !ok {
				break loop
			}
			sizes[resp.root].nbytes += resp.nbytes
			sizes[resp.root].nfiles++
		case <-tick:
			printDiskUsageForRoots(roots, sizes)
		}
	}
	printDiskUsageForRoots(roots, sizes)
}

func printDiskUsageForRoots(roots []string, sizes []SizeCount) {
	for i, root := range roots {
		fmt.Printf("%s ", root)
		printDiskUsage(sizes[i].nfiles, sizes[i].nbytes)
	}
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d файлов %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
