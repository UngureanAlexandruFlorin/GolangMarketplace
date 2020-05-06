package middlewares

import (
	"fmt"
	"strings"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"github.com/dgrijalva/jwt-go"
	"local.com/golangMarketplace/productService/models"
)

var jwtKey = []byte("secretJwtKey")

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(responseWriter http.ResponseWriter, request *http.Request) {
			enableCors(&responseWriter);

			if (decodeJWT(responseWriter, request)) {
				next.ServeHTTP(responseWriter, request);
			}
	});
}

func decodeJWT(responseWriter http.ResponseWriter, request *http.Request) bool {

	if (
		len(request.Header["Authorization"]) == 0 ||
		!strings.Contains(request.Header["Authorization"][0], " ")) {
		responseWriter.WriteHeader(http.StatusUnauthorized);
		fmt.Fprintf(responseWriter, "Missing headers!");
		return false;
	}

	var jwtToken []string = strings.Split(request.Header["Authorization"][0], " ");
	var success = true;

	if (len(jwtToken) != 2) {
		responseWriter.WriteHeader(http.StatusUnauthorized);
		fmt.Fprintf(responseWriter, "Invalid token!");
		return false;
	}

	var tokenString string = jwtToken[1];

	claims := &models.JwtClaims{};

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	});

	if (err != nil) {
		if (err == jwt.ErrSignatureInvalid) {
			responseWriter.WriteHeader(http.StatusUnauthorized);
			fmt.Fprintf(responseWriter, err.Error());
			return false;
		}

		responseWriter.WriteHeader(http.StatusBadRequest);
		fmt.Fprintf(responseWriter, err.Error());
		return false;
	}

	if (!token.Valid) {
		responseWriter.WriteHeader(http.StatusUnauthorized);
		fmt.Println("not valid");
		return false;
	}

	var decodedBody map[string]interface{};

	json.NewDecoder(request.Body).Decode(&decodedBody);
	decodedBody["jwtEmail"] = claims.Email;

	jsonBody, err := json.Marshal(decodedBody);

	request.Body = ioutil.NopCloser(strings.NewReader(string(jsonBody)));

	return success;
}

func enableCors(w *http.ResponseWriter) {
    (*w).Header().Set("Access-Control-Allow-Origin", "*");
    (*w).Header().Set("Content-Type", "application/json");
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization");
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE");
    (*w).Header().Set("Access-Control-Allow-Credentials", "true");
}