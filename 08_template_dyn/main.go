package main

import (
	"log"
	"net/http"
	"text/template"
)

type Context struct {
	FirstName string
	Mess      string
}

func main() {
	http.HandleFunc("/", someFunc)
	http.HandleFunc("/Hubert", HubertFunc)
	http.HandleFunc("/Wolf", WolfFunc)
	http.ListenAndServe(":8080", nil)
}

func someFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	tmpl, err := template.New("AnyName").Parse(doc)
	log.Println(req.URL.Path)
	if err == nil {
		context := Context{"Todd", "more Go, please!"}
		tmpl.Execute(w, context)
	}
}

func HubertFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	tmpl, err := template.New("AnyName").Parse(doc)
	log.Println(req.URL.Path)
	if err == nil {
		context := Context{"Hubert", "I'm trouble maker"}
		tmpl.Execute(w, context)
	}
}

func WolfFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	tmpl, err := template.New("AnyName").Parse(doc)
	log.Println(req.URL.Path)
	if err == nil {
		context := Context{"Mr Wolf", "I solve problems"}
		tmpl.Execute(w, context)
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
</body>
</html>
`
