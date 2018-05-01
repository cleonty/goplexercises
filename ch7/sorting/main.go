package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

type byArtist []*Track

func (s byArtist) Len() int {
	return len(s)
}

func (s byArtist) Less(i, j int) bool {
	return s[i].Artist < s[j].Artist
}

func (s byArtist) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type reverse struct {
	sort.Interface
}

func (s reverse) Less(i, j int) bool {
	return s.Interface.Less(j, i)
}

func Reverse(data sort.Interface) sort.Interface {
	return reverse{data}
}

type customSort struct {
	t    []*Track
	less func(i, j *Track) bool
}

func (s customSort) Len() int {
	return len(s.t)
}

func (s customSort) Less(i, j int) bool {
	return s.less(s.t[i], s.t[j])
}

func (s customSort) Swap(i, j int) {
	s.t[i], s.t[j] = s.t[j], s.t[i]
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "----", "------", "-----", "---", "-----")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

func main() {
	sort.Sort(sort.Reverse(byArtist(tracks)))
	sort.Sort(Reverse(byArtist(tracks)))
	sort.Sort(customSort{tracks, func(i, j *Track) bool {
		if i.Title != j.Title {
			return i.Title < j.Title
		}
		if i.Year != j.Year {
			return i.Year < j.Year
		}
		if i.Length != j.Length {
			return i.Length < j.Length
		}
		return false
	}})
	printTracks(tracks)
	values := []int{3, 1, 4, 2}
	fmt.Println(sort.IntsAreSorted(values))
	sort.Ints(values)
	fmt.Println(sort.IntsAreSorted(values))
	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	fmt.Println(values)
	fmt.Println(sort.IntsAreSorted(values))
}
