package middlewares

import (
	"net/http"
	"local.com/golangMarketplace/productService/controllers"
)

func InitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(responseWriter http.ResponseWriter, request *http.Request) {
			controllers.Init();
			next.ServeHTTP(responseWriter, request);
	});
}