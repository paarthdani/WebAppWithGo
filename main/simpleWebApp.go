package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	new := "Variable"
	fmt.Fprintf(w, "<h1>This is a Web App!</h1>")
	fmt.Fprintf(w, "<p>which was made using GoLang</p>")
	fmt.Fprintf(w, "<p>GoLang is fast and simple.</p>")
	fmt.Fprintf(w, `
	<p> this is a multiline print.
	I am passing a variable here - %s
	`, new)
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is an about page designed by Parth!")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about", about_handler)
	http.ListenAndServe(":8000", nil)
}
