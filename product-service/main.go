package main

import (
	"fmt"
	"net/http"
	"local.com/golangMarketplace/productService/controllers"
	"local.com/golangMarketplace/productService/middlewares"
)

func main() {
	/*
		I know there is something like a server mixin that would allow
		me to declare one root path (like "/") and then go to sub routes
		allowing me to apply the middlewares only on the root path.

		I will refactor this later.
	 */

	fmt.Printf("Product service started on port 8082!\n");
	var createHandler http.Handler = http.HandlerFunc(controllers.Create);
	var getAllHandler http.Handler = http.HandlerFunc(controllers.GetAll);
	var getByIdHandler http.Handler = http.HandlerFunc(controllers.GetById);
	var getByEmailHandler http.Handler = http.HandlerFunc(controllers.GetByEmail);
	var updateHandler http.Handler = http.HandlerFunc(controllers.Update);
	var deleteHandler http.Handler = http.HandlerFunc(controllers.Delete);

	http.Handle("/create", middlewares.AuthMiddleware(middlewares.InitMiddleware(createHandler)));
	http.Handle("/get", middlewares.AuthMiddleware(middlewares.InitMiddleware(getAllHandler)));
	http.Handle("/getById", middlewares.AuthMiddleware(middlewares.InitMiddleware(getByIdHandler)));
	http.Handle("/getByEmail", middlewares.AuthMiddleware(middlewares.InitMiddleware(getByEmailHandler)));
	http.Handle("/update", middlewares.AuthMiddleware(middlewares.InitMiddleware(updateHandler)));
	http.Handle("/delete", middlewares.AuthMiddleware(middlewares.InitMiddleware(deleteHandler)));
	http.ListenAndServe(":8082", nil);
}