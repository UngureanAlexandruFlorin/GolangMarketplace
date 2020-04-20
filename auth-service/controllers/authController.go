package controllers

import (
    "fmt"
    "net/http"
    "github.com/dgrijalva/jwt-go"
    "encoding/json"
    "time"
    "database/sql"
	_ "github.com/lib/pq"
)

var jwtSecret []byte = []byte("secretJwtKey");
var db *sql.DB;
var err error;

func check(err error) {
    if err != nil {
        panic(err)
    }
}

type User struct {
    Email string;
    Password string;
}

func Login(responseWriter http.ResponseWriter, request *http.Request) {
	var user User;

    check(json.NewDecoder(request.Body).Decode(&user));

    if (db == nil) {
    	connectToBD();
    }

    rows, error := db.Query(`select email, password from users where email = $1 and password = $2;`, user.Email, user.Password);
	check(error);

	if (!rows.Next()) {
		responseWriter.WriteHeader(http.StatusUnauthorized);
		fmt.Printf("Error! User don't exist! Email: %s\n", user.Email);
		fmt.Fprintf(responseWriter, "Error! User doesn't exist! Email: %s\n", user.Email);
	}

    token := jwt.New(jwt.SigningMethodHS256);

    claims := token.Claims.(jwt.MapClaims);
    claims["authorized"] = true;
    claims["email"] = user.Email;
    claims["exp"] = time.Now().Add(time.Minute * 30).Unix();

    tokenString, error := token.SignedString(jwtSecret);
    check(error);

    fmt.Fprintf(responseWriter, "Bearer %s", tokenString);
}

func Register(responseWriter http.ResponseWriter, request *http.Request) {
	var user User;

    check(json.NewDecoder(request.Body).Decode(&user));

    if (db == nil) {
    	connectToBD();
    }

	rows, error := db.Query(`select email from users where email = $1;`, user.Email);
	check(error);

	if (rows.Next()) {
		responseWriter.WriteHeader(http.StatusBadRequest);
		fmt.Printf("Error! User already exist! Email: %s\n", user.Email);
		fmt.Fprintf(responseWriter, "Error! User already exist! Email %s\n", user.Email);
		return;
	}

	rows, error = db.Query(`insert into users (email, password) values ($1, $2);`, user.Email, user.Password);
	check(error);

    token := jwt.New(jwt.SigningMethodHS256);

    claims := token.Claims.(jwt.MapClaims);
    claims["authorized"] = true;
    claims["email"] = user.Email;
    claims["exp"] = time.Now().Add(time.Minute * 30).Unix();

    tokenString, error := token.SignedString(jwtSecret);
    check(error);

    fmt.Fprintf(responseWriter, "Bearer %s", tokenString);
}

func connectToBD() {
	fmt.Println("Start auth module!");
	connStr := "postgres://ec2-user:password@172.31.0.5/golang_marketplace?sslmode=disable"; // I know this should be inside a config file. I will make that later.
	db, err = sql.Open("postgres", connStr);
	check(err);
}