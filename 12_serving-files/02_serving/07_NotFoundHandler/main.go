package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", bird)
	http.Handle("/favi.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func bird(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL)
	fmt.Fprintln(w, "go look at terminal")
}
