package main

import (
	"os"
)

func tempDirs() []string {
	return []string{"1", "2", "3"}
}

func main() {
	var rmdirs []func()
	for _, d := range tempDirs() {
		dir := d
		os.Mkdir(dir, 0755)
		rmdirs = append(rmdirs, func() { os.Remove(dir) })
	}
	for _, rmdir := range rmdirs {
		rmdir()
	}
}
