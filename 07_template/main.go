package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", someFunc)
	http.ListenAndServe(":8080", nil)
}

func someFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	tmpl, err := template.New("AnyName").Parse(doc)
	log.Println(req.URL.Path)
	if err == nil {
		tmpl.Execute(w, req.URL.Path[1:])
	}
}

const doc = `
<!DOCTYPE html>
<html>
<head lang="en">
	<title>First Template</title>

</head>
<body>
	<h1>Hello {{.}}!</h1>
</body>
</html>
`
