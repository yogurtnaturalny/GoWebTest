package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type MyHandler struct {
}

func (*MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	log.Println(path)

	data, err := ioutil.ReadFile(string(path))

	if err == nil {
		var contentType string

		if strings.HasSuffix(path, ".css") {
			contentType = "text/css"

		} else if strings.HasSuffix(path, ".html") {
			contentType = "text/html"

		} else if strings.HasSuffix(path, ".js") {
			contentType = "application/javascript"

		} else if strings.HasSuffix(path, ".png") {
			contentType = "image/png"

		} else {
			contentType = "text/plain"

		}

		w.Header().Add("Content-Type", contentType)
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 for you this time!" + http.StatusText(404)))
	}

}

func main() {
	http.Handle("/", new(MyHandler))
	http.ListenAndServe(":8080", nil)
}
