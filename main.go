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

func contact(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("public/contact.html"))
	tmpl.Execute(w, nil)
}

type Client struct {
	Name      string
	Firstname string
	Phone     string
}

func traitement(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("public/traitement.html"))
	if r.Method != "POST" {
		tmpl.Execute(w, nil)
		return
	}

	data := Client{
		Name:      r.FormValue("name"),
		Firstname: r.FormValue("fName"),
		Phone:     r.FormValue("phone"),
	}

	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/traitement", traitement)

	http.ListenAndServe(":3000", nil)
}
