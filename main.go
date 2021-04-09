package main

import (
	"fmt"
	"net/http"
)

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/", PrintGreeting)
	http.ListenAndServe("0.0.0.0:8080", handler)
}

func PrintGreeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"message": "Hello world"}`)
}
