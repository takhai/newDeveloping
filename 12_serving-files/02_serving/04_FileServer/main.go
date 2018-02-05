package main

import (
	"io"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/bird", bird)
	http.ListenAndServe(":8080", nil)
}

func bird(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=itf-8")
	io.WriteString(w, `<img src="blaky.jpg">`)
}
