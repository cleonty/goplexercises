package main

import (
	"log"
	"os"
	"text/template"
	"time"

	"github.com/cleonty/gopl/ch4/github"
)

const templ = `{{.TotalCount}} тем
{{range .Items}}----------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}
`

var report = template.Must(template.New("report").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func main() {
	var terms []string = []string{"json"}
	if len(os.Args) > 1 {
		terms = os.Args[1:]
	}
	result, err := github.SearchIssues(terms)
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}

}
