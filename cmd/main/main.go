package main

import (
	"chatserver/pkg/db"
	"chatserver/pkg/routes"
	"fmt"
	"net/http"
)

func main() {
	dbPool, err := db.InitializeDB()
	if err != nil {
		fmt.Printf("Error: %v ", err)
	}
	defer dbPool.Close()

	http.HandleFunc("/register", routes.Register)
	http.HandleFunc("/login", routes.Login)
	http.ListenAndServe(":8080", nil)
}
