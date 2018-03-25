package main

import (
	"log"
	"net/http"
	"text/template"
)

type context struct {
	FirstName string
	Mess      string
	URL       string
	Beers     []string
	Title     string
}

func main() {
	http.HandleFunc("/", someFunc)
	http.ListenAndServe(":8080", nil)
}

func someFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	templates := template.New("template")

	log.Println(req.URL.Path)
	templates.New("test").Parse(doc)
	templates.New("header").Parse(head)
	templates.New("footer").Parse(foot)
	Context := context{
		"Hubert",
		"More beer, please!",
		req.URL.Path,
		[]string{"Tyskie", "Tesco Value", "Tesco Finest"},
		"Awsome beers",
	}
	templates.Lookup("test").Execute(w, Context)
}

const doc = `
{{template "header" .Title}}

<body>
	<h1>Hello {{.FirstName}}!</h1>
	<p>{{.Mess}}</p>

{{if eq .URL "/nobeer"}}
	<h2>Out of beer, {{.FirstName}}</h2>
{{else}}
	<h2>Yes, let's have some beer, {{.FirstName}}</h2>
	<ul>
		{{range .Beers}}
		<li>{{.}}</li>
		{{end}}
	</ul>
{{end}}
<hr>

<h2>Here is some data:</h2>
<p>{{.}}</p>
</body>

{{template "footer"}}
`

const head = `
<!DOCTYPE html>
<html>
<head lang="en">
	<title>First Template</title>

</head>
`

const foot = `
</html>
`
