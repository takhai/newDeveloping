package main

import (
	"github.com/satori/go.uuid"
	"html/template"
	"net/http"
)

//check code

type user struct {
	UserName string
	First    string
	Last     string
}

var tpl *template.Template
var dbUsers = map[string]user{}      //user iD , user
var dbSessions = map[string]string{} //session ID

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/process", bar)
	http.Handle("/favi.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foe(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func process(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	fname := r.FormValue("firster")
	lname := r.FormValue("laster")

	d := struct {
		First string
		Last  string
	}{
		First: fname,
		Last:  lname,
	}
	tpl.ExecuteTemplate(w, "processor.gohtml", d)

}
