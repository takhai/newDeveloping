package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.Handle("/assets", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.Handle("/favi.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", "MO inc")
}

func about(w http.ResponseWriter, r *http.Request) {
	type customData struct {
		Title   string
		Members []string
	}

	cd := customData{
		Title:   "The Team",
		Members: []string{"Moshushu", "Takhai", "Maisja", "Ameena"},
	}
	tpl.ExecuteTemplate(w, "about.gohtml", cd)
}
