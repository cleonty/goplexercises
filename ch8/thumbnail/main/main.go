package main

import (
	"log"
	"os"
	"sync"

	"github.com/cleonty/gopl/ch8/thumbnail"
)

func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		if _, err := thumbnail.ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		go thumbnail.ImageFile(f)
	}
}

func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(filename string) {
			thumbnail.ImageFile(filename)
			ch <- struct{}{}
		}(f)
	}
	for range filenames {
		log.Println("finished with one file")
		<-ch
	}
}

func makeThumbnails4(filenames []string) error {
	errors := make(chan error)
	for _, f := range filenames {
		go func(filename string) {
			_, err := thumbnail.ImageFile(filename)
			if err != nil {
				errors <- err
			}
		}(f)
	}
	for range filenames {
		if err := <-errors; err != nil {
			log.Println("ошибка")
			return err
		}
	}
	return nil
}

func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}
	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(filename string) {
			var it item
			it.thumbfile, it.err = thumbnail.ImageFile(filename)
			ch <- it
		}(f)
	}
	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}
	return thumbfiles, nil
}

func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup
	for f := range filenames {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			thumbfile, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumbfile)
			sizes <- info.Size()
		}(f)
	}
	go func() {
		wg.Wait()
		close(sizes)
	}()
	var total int64
	for size := range sizes {
		total += size
	}
	return total
}

func main() {
	files := []string{"./0.jpg", "./1.jpg", "./2.jpg", "./3.jpg", "./4.jpg"}
	ch := make(chan string)
	go func() {
		for _, f := range files {
			ch <- f
		}
		close(ch)
	}()
	total := makeThumbnails6(ch)
	log.Println("done")
	log.Println(total)
}
