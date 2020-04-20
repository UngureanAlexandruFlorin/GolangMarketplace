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

    fmt.Printf("User data: email %s - password %s\n", user.Email, user.Password);

	rows, error := db.Query(`select email, password from users where email = $1;`, user.Email);
	check(error);

	var result string = "";

	for rows.Next() {
		var email string;
		var password string;

		error := rows.Scan(&email, &password);
		check(error);

		result = result + email + " - " + password + "\n"
	}

    // token := jwt.New(jwt.SigningMethodHS256);

    // claims := token.Claims.(jwt.MapClaims);
    // claims["authorized"] = true;
    // claims["email"] = user.email;
    // claims["exp"] = time.Now().Add(time.Minute * 30).Unix();

    // tokenString, error := token.SignedString(jwtSecret);
    // check(error);


    // fmt.Fprintf(responseWriter, "Bearer %s", tokenString);
    fmt.Fprintf(responseWriter, "Register result: %s\n", result);
}

func connectToBD() {
	fmt.Println("Start auth module!");
	connStr := "postgres://ec2-user:password@172.31.0.5/golang_marketplace?sslmode=disable"; // I know this should be inside a config file. I will make that later.
	db, err = sql.Open("postgres", connStr);
	check(err);
}