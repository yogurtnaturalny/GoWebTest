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
	//Beers     []string
	//Title     string
}

func main() {
	http.HandleFunc("/", someFunc)
	http.ListenAndServe(":8080", nil)
}

func someFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	tmpl, err := template.New("AnyName").Parse(doc)
	log.Println(req.URL.Path)
	if err == nil {
		Context := context{
			"Hubert",
			"more beer, please",
			req.URL.Path,
			//	[]string{"Tyskie", "Tesco Value", "Tesco Finest"},
			//	"Awsome beers",
		}
		tmpl.Execute(w, Context)
	} else {
		log.Panic("Something go wrong!")

	}
}

const doc = `
<!DOCTYPE html>
<html>
<head lang="en">
	<title>First Template</title>

</head>
<body>
	<h1>Hello {{.FirstName}}!</h1>
	<p>{{.Mess}}</p>

{{if eq .URL "/nobeer"}}
	<h2>Out of beer, {{.FirstName}}</h2>
{{else}}
	<h2>Yes, let's have some beer, {{.FirstName}}

{{end}}
<hr>

<h2>Here is some data:</h2>
<p>{{.URL}}</p>
</body>
</html>
`
