package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	buddha := sage{
		Name:  "Buddha",
		Motto: "Life is good so live good",
	}

	jesus := sage{
		Name:  "Jesus",
		Motto: "Drink wine",
	}

	krishna := sage{
		Name:  "Krishna",
		Motto: "Have good sex",
	}

	sages := []sage{buddha, jesus, krishna}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
