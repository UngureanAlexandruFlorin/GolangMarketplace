package controllers

import (
	"context"
	"fmt"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"local.com/golangMarketplace/productService/models"
)

var ctx context;
var client *mongo.Client;
var products *mongo.Collection;

func Create(responseWriter http.ResponseWriter, request *http.Request) {
	var product models.Product;

	check(json.NewDecoder(request.Body).Decode(&product));

	insertResult, err := products.InsertOne(ctx, bson.D {
		{ Key: "email", Value: "test@gmail.com" },
		{ Key: "password", Value: "testPass" },
	});

	check(err);

	fmt.Fprintf(responseWriter, "-");

}

func Read(responseWriter http.ResponseWriter, request *http.Request) {
	var product models.Product;

	err := products.FindOne(ctx, bson.M{"password": "testPass"}).Decode(&product);

	check(err);

	fmt.Fprintf(responseWriter, product.Email);
}

func Update(responseWriter http.ResponseWriter, request *http.Request) {

}

func Delete(responseWriter http.ResponseWriter, request *http.Request) {

}

func init() {
	var err error;

	ctx, _ = context.WithTimeout(context.Background(), 10 * time.Second);
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"));
	err = client.Connect(ctx);

	if (err != nil) {
		panic(err);
	}

	defer client.Disconnect(ctx);

	database := client.Database("golang_marketplace");

	products = database.Collection("products");

	// if err = cursor.All(ctx, &episodes); err != nil {
 //    	panic(err);
	// }
}

func check(err error) {
	if (err != nil) {
		panic(err);
	}
}