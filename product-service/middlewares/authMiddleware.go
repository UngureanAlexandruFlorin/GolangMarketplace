package middlewares

import (
	"fmt"
	"strings"
	"net/http"
	"github.com/dgrijalva/jwt-go"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(responseWriter http.ResponseWriter, request *http.Request) {
			decodeJWT(responseWriter, request);
			next.ServeHTTP(responseWriter, request);
	});
}

func decodeJWT(responseWriter http.ResponseWriter, request *http.Request) {
	var jwtToken []string = strings.Split(request.Header["Authorization"][0], " ");

	if (len(jwtToken) != 2) {
		fmt.Fprintf(responseWriter, "Invalid token!");
		panic("Invalid token!");
	}

	var tokenString string = jwtToken[1];

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
        	return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
    	}
    	return "secret", nil
	});

	if (err != nil) {
		panic(err);
	}

	claims, ok := token.Claims.(jwt.MapClaims);

	if (!ok) {
		fmt.Fprintf(responseWriter, "Failed to claim data from JWT token!");
		panic("Failed to claim data from JWT token!");
	}

	if (ok && token.Valid) {
		fmt.Println(claims["email"], claims["exp"]);
	}
}