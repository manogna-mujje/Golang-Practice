package main

import (
	"net/http"
	"html/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w, "index.gohtml", "WEBSITE MODEL")
}


func about(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w, "about.gohtml", "ABOUT US")
}
