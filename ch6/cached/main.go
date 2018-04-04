package main

import (
	"fmt"
	"sync"
)

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func Lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}

func Add(key, value string) {
	cache.Lock()
	cache.mapping[key] = value
	cache.Unlock()
}

func main() {
	Add("123", "456")
	fmt.Println(Lookup("123"))
}
