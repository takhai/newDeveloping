package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", bird)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}

func bird(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=itf-8")
	io.WriteString(w, `<img src="/resources/blaky.jpg">`)
}
