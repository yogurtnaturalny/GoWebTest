package main

import (
	"net/http"
)

type person struct {
	fName string
}

func (p *person) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("First Name: " + p.fName))
}

func main() {
	personOne := &person{fName: "Hubert"}
	http.ListenAndServe(":8080", personOne)
}
