package main

import (
	"html/template"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("public/index.html"))
	tmpl.Execute(w, nil)
}

func about(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("public/about.html"))
	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)

	http.ListenAndServe(":3000", nil)
}
