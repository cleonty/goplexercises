package main

import (
	"html/template"
	"log"
	"os"
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

var issueList = template.Must(template.New("issuelist").Parse(`<h1>{{.TotalCount}} Teм</h1>
<table>
  <tr style='text-align: left'>
    <th>#</th>
    <th>State</th>
    <th>User</th>
    <th>Title</th>
  </tr>
  {{range .Items}}
  <tr>
    <td><a href='{{.HTMLURL}}'>{{.Number}}</a>/td>
    <td>{{.State}}</td>
    <td><a href='{{.User.HTMLURL}}'>{{ .User.Login}}</a></td> <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
  </tr>
  {{end}}
</table>
`))

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
	if err := issueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}

}
