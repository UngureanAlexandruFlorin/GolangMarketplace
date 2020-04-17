package main

import (
    "fmt"
    "net/http"
    // "local.com/golangMarketplace/authService/controllers"
)

func loginHandler(responseWriter http.ResponseWriter, request *http.Request) {
    // controllers.Login(responseWriter, request);
    fmt.Fprintf(responseWriter, "login");
    fmt.Println("login");
}

func registerHandler(responseWriter http.ResponseWriter, request * http.Request){
    // controllers.Register(responseWriter, request);
    fmt.Fprintf(responseWriter, "register");
    fmt.Println("register");
}

func main() {
    http.HandleFunc("/login", loginHandler);
    http.HandleFunc("/register", registerHandler);
    http.ListenAndServe(":8080", nil);

    fmt.Printf("Server started on port 8080!");

}