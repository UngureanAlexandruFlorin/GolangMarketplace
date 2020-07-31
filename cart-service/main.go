package main

import (
	"fmt"

	"local.com/golangMarketplace/cartService/routes"
)

func main() {
	fmt.Println("Cart service started on port 8083!")

	var router *routes.Router = new(routes.Router)
	router.Use("/cart", routes.RoutesRouter)

	router.ListenAndServe(":8083", nil)
}
