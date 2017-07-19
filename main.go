package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "I guess we can be cliche and use the all too popular - Hello World")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":80", nil)
}