package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "app2")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9002", nil)
}