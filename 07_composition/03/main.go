package main

import (
	"log"
	"os"
	"text/template"
)

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	y := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				course{"DT-01", "Introduction", "100"},
				course{"DT-02", "Bit to Hashgraph", "200"},
				course{"DT-03", "Hashgraph to GhostMap", "300"},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				course{"DT-04", "Ai and EI", "075"},
				course{"DT-05", "Space VS Virtual Space", "145"},
				course{"DT-06", "LLC neurospeed tracking", "345"},
			},
		},
	}

	err := tpl.Execute(os.Stdout, y)
	if err != nil {
		log.Fatalln(err)
	}
}
