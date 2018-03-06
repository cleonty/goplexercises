package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func download(num int) error {
	url := "https://xkcd.com/" + strconv.Itoa(num) + "/info.0.json"
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	var body []byte
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	filename := "./data/" + strconv.Itoa(num) + ".json"
	err = ioutil.WriteFile(filename, body, 0644)
	if err != nil {
		return err
	}
	return nil
}

type Comics struct {
	Num        int
	Title      string `json:"safe_title"`
	Transcript string
}

func read(num int) (*Comics, error) {
	filename := "./data/" + strconv.Itoa(num) + ".json"
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var comics Comics
	err = json.NewDecoder(file).Decode(&comics)
	if err != nil {
		return nil, err
	}
	return &comics, nil
}

func match(terms []string, comics *Comics) bool {
	termCount := len(terms)
	matchCount := 0
	for _, term := range terms {
		if strings.Contains(comics.Title, term) {
			matchCount++
		}
	}
	return matchCount == termCount
}

func search(terms []string) {
	for i := 1; i < 1000; i++ {
		comics, err := read(i)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if match(terms, comics) {
			fmt.Printf("%d %s [%s]\n", comics.Num, comics.Title, comics.Transcript)
		}

	}
}

func downloadAll() {
	for i := 1; i < 1000; i++ {
		fmt.Println("Downloading file", i)
		err := download(i)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func main() {
	search(os.Args[1:])
}
