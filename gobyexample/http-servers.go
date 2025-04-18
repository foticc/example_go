package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	fmt.Fprintln(w, "Hello, world!")
}

func headers(w http.ResponseWriter, r *http.Request) {
	for name, value := range r.Header {
		fmt.Fprintf(w, "%s: %s\n", name, value)
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/", headers)

	http.ListenAndServe(":8000", nil)
}
