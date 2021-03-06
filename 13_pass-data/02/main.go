package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favi.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) { //fix

	f := req.FormValue("first")
	//l := req.FormValue("last")
	//s := req.FormValue("subscribe") == "on"

	tpl.ExecuteTemplate(w, "index.gohtml", f)
}
