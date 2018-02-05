package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favi.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	w.Header().Set("Content-type", "text/html; charset=itf-8")
	io.WriteString(w, `
		<form method="POST">
			<input type="text" name="q">
			<input type="submit">
		</form>
		<br>`+v)
}

//if you change METHOD to GET it will show it in the URL
