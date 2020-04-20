package main

import (
    "fmt"
    "net/http"
    "local.com/golangMarketplace/authService/controllers"
)

func main() {
    fmt.Printf("Server started on port 8080!");
    
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/register", controllers.Register)
	http.ListenAndServe(":8080", nil)
}
