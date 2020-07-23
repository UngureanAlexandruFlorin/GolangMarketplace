package main

import (
	"fmt"
	"net/http"

	"local.com/golangMarketplace/authService/controllers"
)

func main() {
	fmt.Printf("Auth service started on port 8081!\n")

	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/register", controllers.Register)
	http.ListenAndServe(":8081", nil)
}
