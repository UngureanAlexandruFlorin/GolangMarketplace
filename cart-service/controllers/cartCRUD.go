package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"local.com/golangMarketplace/cartService/models"
)

var ctx context.Context
var client *mongo.Client
var cart *mongo.Collection

func Create(res http.ResponseWriter, req *http.Request) {
	var err error
	var existingCart models.Cart
	var newCartProduct models.Product

	err = json.NewDecoder(req.Body).Decode(&newCartProduct)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(res, "Invalid body data!")
		fmt.Println("Invalid body data!\n" + err.Error())
		return
	}

	err = cart.FindOne(ctx, bson.M{"email": newCartProduct.JwtEmail}).Decode(&existingCart)

	if err != nil && err != mongo.ErrNoDocuments {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(res, "Invalid body data!")
		fmt.Println("Invalid body data!\n" + err.Error())
		return
	}

	if err == mongo.ErrNoDocuments {
		_, err = cart.InsertOne(ctx, bson.M{
			"email": newCartProduct.JwtEmail,
			"products": bson.A{
				bson.M{
					"name":        newCartProduct.Name,
					"description": newCartProduct.Description,
					"price":       newCartProduct.Price,
				},
			},
		})
	} else {
		_, err = cart.UpdateOne(ctx, bson.M{
			"email": newCartProduct.JwtEmail,
		},
			bson.M{"products": bson.D{{
				"$push", bson.D{
					{"name", newCartProduct.Name},
					{"description", newCartProduct.Description},
					{"price", newCartProduct.Price},
				},
			}},
			})
	}

	fmt.Fprintf(res, "Done")
}

func Read(res http.ResponseWriter, req *http.Request) {
	var err error
	var existingCart models.Cart

	var jwtEmail models.GetJwtEmail

	err = json.NewDecoder(req.Body).Decode(&jwtEmail)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(res, "Invalid Auth header!")
		fmt.Println("Invalid Auth header!\n" + err.Error())
		return
	}

	err = cart.FindOne(ctx, bson.M{"email": jwtEmail.JwtEmail}).Decode(&existingCart)

	if err != mongo.ErrNoDocuments {
		fmt.Fprintf(res, "{}")
		return
	}

	var cartString []byte
	cartString, err = json.Marshal(existingCart)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(res, "Failed to create a response!")
		fmt.Println("Failed to create a response!!\n" + err.Error())
		return
	}

	fmt.Fprintf(res, string(cartString))
}

func Update(res http.ResponseWriter, req *http.Request) {

}

func Delete(res http.ResponseWriter, req *http.Request) {

}

func init() {
	var err error

	ctx = context.Background()
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://admin:password@192.168.1.13:27017"))

	if err != nil {
		panic(err)
	}

	err = client.Connect(ctx)

	if err != nil {
		panic(err)
	}

	database := client.Database("golang_marketplace")
	cart = database.Collection("carts")
}
