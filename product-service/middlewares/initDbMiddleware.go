package middlewares

import (
	"net/http"
)

func InitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(responseWriter http.ResponseWriter, request *http.Request) {
			next.ServeHTTP(responseWriter, request)
		})
}
