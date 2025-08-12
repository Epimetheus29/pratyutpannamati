package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Project Pratyutpannamati!")
}

func main() {
	http.HandleFunc("/", homeHandler)
	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":3000", nil)
}
