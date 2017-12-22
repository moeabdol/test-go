package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer((http.Dir("public")))))
	http.ListenAndServe(":9080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	n := "James Bond"
	tpl.ExecuteTemplate(w, "index.gohtml", n)
}
