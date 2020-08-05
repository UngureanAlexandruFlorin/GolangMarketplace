package models

import "github.com/dgrijalva/jwt-go"

type Cart struct {
	Email    string
	Products []Product
}

type Product struct {
	JwtEmail    string
	ID          string
	Email       string
	Name        string
	Description string
	Price       int
}

type JwtClaims struct {
	Email string
	jwt.StandardClaims
}

type AuthJwtData struct {
	AuthEmail string
	Body      interface{}
}

type GetJwtEmail struct {
	JwtEmail string
}
