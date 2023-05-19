package main

import (
	"html/template"
	"net/http"
)

type FormData struct {
	Name  string
	Email string
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/submit", submitHandler)
	http.ListenAndServe("localhost:8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles("index.html"))
	tmp.Execute(w, nil)
}
