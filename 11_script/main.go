package main

import (
	"html/template"
	"log"
	"net/http"
	//"text/template"
)

//var Message string = "More beer, please!"

var Message string = "alert('you have been pwned')"

//var Message string = "<script>alert('you have been pwned, man!')</script>"

func main() {
	http.HandleFunc("/", someFunc)
	http.ListenAndServe(":8080", nil)
}

func someFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	tmpl, err := template.New("AnyName").Parse(doc)
	log.Println(req.URL.Path)
	if err == nil {
		tmpl.Execute(w, Message)
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

	<p>{{.}}</p>

	<script>{{.}}</script>

</body>
</html>
`
