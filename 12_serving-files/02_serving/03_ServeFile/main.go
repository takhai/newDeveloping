package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/blaky.jpg", birdPic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-type", "text/html; charset=itf-8")

	io.WriteString(w, `
		<img src="blaky.jpg">
		`)
}

func birdPic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "blaky.jpg")
}
