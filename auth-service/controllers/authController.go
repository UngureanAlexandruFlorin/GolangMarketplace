package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

var jwtSecret []byte = []byte("secretJwtKey")

func check(error error) {
	if error != nil {
		panic(error)
	}
}

type User struct {
	email    string
	password string
}

func Login(responseWriter http.ResponseWriter, request *http.Request) {
	var user User

	check(json.NewDecoder(request.Body).Decode(&user))

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["email"] = user.email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, error := token.SignedString(jwtSecret)
	check(error)

	fmt.Fprintf(responseWriter, "Bearer %s", tokenString)
}

func Register(responseWriter http.ResponseWriter, request *http.Request) {
	var user User

	check(json.NewDecoder(request.Body).Decode(&user))

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["email"] = user.email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, error := token.SignedString(jwtSecret)
	check(error)

	fmt.Fprintf(responseWriter, "Bearer %s", tokenString)
}
