package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favi.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "tak-cookie",
		Value: "alot",
	})
	fmt.Fprintln(w, "Cookie written-check rowser")
	fmt.Fprintln(w, "go to dev tools")
}

func read(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("tak-cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, "Your Cookie:", c)
}
