package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-type", "text/html; charset=itf-8")

	io.WriteString(w, `
		<!--not serving from here-->,
		<img src="http://rawalpindi.bolee.com/resources/items/img/2014/01/08/12470928201389164663.jpg"><h1>Pigeon</h1>
		`)
}
