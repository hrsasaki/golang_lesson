// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 115.

// Issueshtml prints an HTML table of issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"html/template"

	"../github"
)

//!+template

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

//!-template

//!+
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	result, err := github.SearchIssues([]string{"repo:golang/go", "json"})
	if err != nil {
		log.Fatal(err)
	}
	if err := issueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, issueList)
}

//!-