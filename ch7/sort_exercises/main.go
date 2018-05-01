package main

import (
	"fmt"
	"html/template"
	"net/http"
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

type dataBase struct {
	t           []*Track
	sortColumns []string
}

func (d dataBase) Len() int {
	return len(d.t)
}

func (d dataBase) Less(i, j int) bool {
	for _, c := range d.sortColumns {
		if c == "title" {
			if d.t[i].Title == d.t[j].Title {
				continue
			} else {
				return d.t[i].Title < d.t[j].Title
			}
		}
		if c == "album" {
			if d.t[i].Album == d.t[j].Album {
				continue
			} else {
				return d.t[i].Album < d.t[j].Album
			}
		}
		if c == "artist" {
			if d.t[i].Artist == d.t[j].Artist {
				continue
			} else {
				return d.t[i].Artist < d.t[j].Artist
			}
		}
		if c == "year" {
			if d.t[i].Year == d.t[j].Year {
				continue
			} else {
				return d.t[i].Year < d.t[j].Year
			}
		}
		if c == "length" {
			if d.t[i].Length == d.t[j].Length {
				continue
			} else {
				return d.t[i].Length < d.t[j].Length
			}
		}
	}
	return false
}

func (d dataBase) Swap(i, j int) {
	d.t[i], d.t[j] = d.t[j], d.t[i]
}

func (d *dataBase) sortBy(column string) {
	d.sortColumns = append([]string{column}, d.sortColumns...)
}

var trackList = template.Must(template.New("tracklist").Parse(`<h1>Трэки</h1>
	<table>
	  <tr>
		<th><a href="?sort=title">Title</a></th>
		<th><a href="?sort=artist">Artist</a></th>
		<th><a href="?sort=album">Album</a></th>
		<th><a href="?sort=year">Year</a></th>
		<th><a href="?sort=length">Length</a></th>
	  </tr>
	  {{range .}}
	  <tr>
	  	<td>{{.Title}}</td>
	  	<td>{{.Artist}}</td>
	  	<td>{{.Album}}</td>
	  	<td>{{.Year}}</td>
	  	<td>{{.Length}}</td>
	  </tr>
	  {{end}}
	</table>
	<a href="?stop=true">Остановить сервер!</a>
	`))

func HandleTracks(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprint(w, err)
		return
	}
	sortColumn := r.Form.Get("sort")
	if len(sortColumn) > 0 {
		ds.sortBy(sortColumn)
		sort.Sort(ds)
	}
	stop := r.Form.Get("stop")
	if len(stop) > 0 {
		os.Exit(0)
	}
	if err := trackList.Execute(w, tracks); err != nil {
		fmt.Fprint(w, err)
		return
	}
}

var ds dataBase = dataBase{tracks, nil}

func main() {
	//ds.sortBy("length")
	//ds.sortBy("year")
	//ds.sortBy("title")
	sort.Sort(ds)
	printTracks(tracks)
	http.HandleFunc("/", HandleTracks)
	http.ListenAndServe(":8080", nil)
}
