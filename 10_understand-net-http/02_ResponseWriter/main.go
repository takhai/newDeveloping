package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Takhai-key", "this is me")
	w.Header().Set("Content-type", "text/html; charset=itf-8")
	fmt.Fprint(w, "<h1>Any Cose you want in ths func</h1>")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
