package main

import (
	"chatserver/pkg/db"
	"fmt"
	"net/http"
)

func registerUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func main() {
	db.Connect()

	http.HandleFunc("/register", registerUser)
	http.ListenAndServe(":8080", nil)
}
