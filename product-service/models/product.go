package models

import "github.com/dgrijalva/jwt-go"

type AlterProduct struct {
	JwtEmail string;
	Id string;
	Email string;
	Name string;
	Description string;
	Price int;
}

type Product struct {
	Email string;
	Name string;
	Description string;
	Price int;
}

type ObjectID struct {
	JwtEmail string;
	Id string;
}

type SellerEmail struct {
	Email string;
}

type JwtClaims struct {
	Email string;
	jwt.StandardClaims;
}

type AuthJwtData struct{
	AuthEmail string;
	Body interface{};
}