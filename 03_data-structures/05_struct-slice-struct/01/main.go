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

type car struct {
	Model string
	Doors int
}

type items struct {
	Wisdom    []sage
	Transport []car
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	b := sage{
		Name:  "Buddha",
		Motto: "Life is good so live good",
	}

	j := sage{
		Name:  "Jesus",
		Motto: "Drink wine",
	}

	k := sage{
		Name:  "Krishna",
		Motto: "Have good sex",
	}

	f := car{
		Model: "Ford",
		Doors: 6,
	}

	v := car{
		Model: "VW",
		Doors: 4,
	}

	p := car{
		Model: "Porche",
		Doors: 2,
	}

	sages := []sage{b, j, k}
	cars := []car{f, v, p}

	data := items{
		Wisdom:    sages,
		Transport: cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
