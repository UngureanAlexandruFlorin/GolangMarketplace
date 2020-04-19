package main

import (
    "fmt"
    "net/http"
    "github.com/local.com/golangMarketplace/authService/controllers"
)

func main() {
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/register", controllers.Register)
	http.ListenAndServe(":8080", nil)

	fmt.Printf("Server started on port 8080!")

}
