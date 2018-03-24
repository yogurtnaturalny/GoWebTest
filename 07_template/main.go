package main

import (
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
	if err == nil {
		tmpl.Execute(w, nil)
	}
}

const doc = `
<!DOCTYPE html>
<html>
<head lang="en">
	<title>First Template</title>

</head>
<body>
	<h1>Hello You!</h1>
</body>
</html>
`
