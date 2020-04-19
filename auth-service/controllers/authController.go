package controllers

import (
    "fmt"
    "net/http"
    "github.com/dgrijalva/jwt-go"
    "encoding/json"
    "time"
)

var jwtSecret []byte = []byte("secretJwtKey");

func check(error error) {
    if error != nil {
        panic(error)
    }
}

type User struct {
    email string;
    password string;
}

func Login(responseWriter http.ResponseWriter, request *http.Request) {
	var user User;

    check(json.NewDecoder(request.Body).Decode(&user));

    token := jwt.New(jwt.SigningMethodHS256);

    claims := token.Claims.(jwt.MapClaims);
    claims["authorized"] = true;
    claims["email"] = user.email;
    claims["exp"] = time.Now().Add(time.Minute * 30).Unix();

    tokenString, error := token.SignedString(jwtSecret);
    check(error);

    fmt.Fprintf(responseWriter, "Bearer %s", tokenString);
}

func Register(responseWriter http.ResponseWriter, request *http.Request) {
	var user User;

    check(json.NewDecoder(request.Body).Decode(&user));

    connStr := "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full"; // I know this should be inside a config file. I will make that later.
	db, err := sql.Open("postgres", connStr);

	rows, err := db.Query(`select * from users where email = $1 `, user.email);

    token := jwt.New(jwt.SigningMethodHS256);

    claims := token.Claims.(jwt.MapClaims);
    claims["authorized"] = true;
    claims["email"] = user.email;
    claims["exp"] = time.Now().Add(time.Minute * 30).Unix();

    tokenString, error := token.SignedString(jwtSecret);
    check(error);

    // fmt.Fprintf(responseWriter, "Bearer %s", tokenString);
    fmt.Fprintf(responseWriter, "Bearer %s", tokenString);
}