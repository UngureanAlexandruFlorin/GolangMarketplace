package main

import (
	"fmt"
	"net/http"
	"local.com/golangMarketplace/productService/controllers"
)

func main() {
	fmt.Printf("Product service started on port 8082!");

	http.HandleFunc("/create", controllers.Create);
	http.HandleFunc("/register", controllers.Read);
	http.ListenAndServe(":8082", nil);
}