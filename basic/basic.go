package main

import "net/http"

func main() {
	myMux := http.NewServeMux()
	myMux.HandleFunc("/", someFunc)
	http.ListenAndServe(":8080", myMux)
}

func someFunc(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello World! v2"))
}
