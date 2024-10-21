package main

import (
	"io"
	"net/http"
)

func setupRouter() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "pong")
	}
}

func main() {
	http.HandleFunc("/ping", setupRouter())

	http.ListenAndServe(":8080", nil)
}
