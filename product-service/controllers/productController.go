package controllers

import (
	"context"
	"fmt"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"local.com/golangMarketplace/productService/models"
)

var ctx context.Context;
var client *mongo.Client;
var products *mongo.Collection;

func Create(responseWriter http.ResponseWriter, request *http.Request) {
	var product models.Product;

	checkWithResponse(json.NewDecoder(request.Body).Decode(&product), responseWriter);

	if (product.Email != product.JwtEmail) {
		responseWriter.WriteHeader(http.StatusUnauthorized);
		fmt.Fprintf(responseWriter, "You can't create a product with someone else's email!");
		return;
	}

	createdProduct, err := products.InsertOne(ctx, bson.D {
		{ Key: "email", Value: product.Email },
		{ Key: "name", Value: product.Name },
		{ Key: "description", Value: product.Description },
		{ Key: "price", Value: product.Price },
	});

	checkWithResponse(err, responseWriter);

	var newObjectId primitive.ObjectID = createdProduct.InsertedID.(primitive.ObjectID);

	fmt.Fprintf(responseWriter, newObjectId.Hex());
	client.Disconnect(ctx);
}

func GetAll(responseWriter http.ResponseWriter, request *http.Request) {
	var foundProducts []models.Product;
	cursor, err := products.Find(ctx, bson.M{});
	checkWithResponse(err, responseWriter);

	for cursor.Next(ctx) {
		var foundProduct models.Product;

		checkWithResponse(cursor.Decode(&foundProduct), responseWriter);
		foundProducts = append(foundProducts, foundProduct);
	}

	jsonResponse, _ := json.Marshal(foundProducts);
	fmt.Fprintf(responseWriter, string(jsonResponse));
	client.Disconnect(ctx);
}

func GetById(responseWriter http.ResponseWriter, request *http.Request) {
	var foundProduct models.Product;
	var id models.ObjectID;

	check(json.NewDecoder(request.Body).Decode(&id));

	objectId, err := primitive.ObjectIDFromHex(id.Id);

	checkWithResponse(err, responseWriter);

	err = products.FindOne(ctx, bson.M{"_id": objectId}).Decode(&foundProduct);

	checkWithResponse(err, responseWriter);

	jsonResponse, _ := json.Marshal(foundProduct);
	fmt.Fprintf(responseWriter, string(jsonResponse));
	client.Disconnect(ctx);
}

func GetByEmail(responseWriter http.ResponseWriter, request *http.Request) {
	var foundProduct models.Product;
	var sellerEmail models.SellerEmail;

	checkWithResponse(json.NewDecoder(request.Body).Decode(&sellerEmail), responseWriter);

	err := products.FindOne(ctx, bson.M{"_id": sellerEmail.Email}).Decode(&foundProduct);
	checkWithResponse(err, responseWriter);

	jsonResponse, _ := json.Marshal(foundProduct);
	fmt.Fprintf(responseWriter, string(jsonResponse));
	client.Disconnect(ctx);
}

func Update(responseWriter http.ResponseWriter, request *http.Request) {
	var newProductData models.Product;
	var foundProduct models.Product;

	checkWithResponse(json.NewDecoder(request.Body).Decode(&newProductData), responseWriter);
	objectId, err := primitive.ObjectIDFromHex(newProductData.Id);
	checkWithResponse(err, responseWriter);

	err = products.FindOne(ctx, bson.M{"_id": objectId}).Decode(&foundProduct);
	checkWithResponse(err, responseWriter);

	if (foundProduct.Email != newProductData.JwtEmail) {
		responseWriter.WriteHeader(http.StatusUnauthorized);
		fmt.Fprintf(responseWriter, "You can't update a product with someone else's email!");
		return;
	}

	updatedObject, err := products.UpdateOne(
    ctx,
    bson.M{"_id": objectId},
    bson.D{
        {"$set", bson.D{
        	{"email", newProductData.Email},
        	{"name", newProductData.Name},
        	{"description", newProductData.Description},
        	{"price", newProductData.Price},
        }},
    });

	jsonResponse, _ := json.Marshal(updatedObject);
    fmt.Fprintf(responseWriter, string(jsonResponse));
    client.Disconnect(ctx);
}

func Delete(responseWriter http.ResponseWriter, request *http.Request) {
	var id models.ObjectID;
	var result *mongo.DeleteResult;
	var foundProduct models.Product;

	checkWithResponse(json.NewDecoder(request.Body).Decode(&id), responseWriter);
	objectId, err := primitive.ObjectIDFromHex(id.Id);
	checkWithResponse(err, responseWriter);

	err = products.FindOne(ctx, bson.M{"_id": objectId}).Decode(&foundProduct);
	checkWithResponse(err, responseWriter);

	if (foundProduct.Email != id.JwtEmail) {
		responseWriter.WriteHeader(http.StatusUnauthorized);
		fmt.Fprintf(responseWriter, "You can't delete a product with someone else's email!");
		return;
	}

	result, err = products.DeleteMany(ctx, bson.M{"_id": objectId})
	
	checkWithResponse(err, responseWriter);

	fmt.Fprintf(responseWriter, "Deleted documents: %d", result.DeletedCount);
	client.Disconnect(ctx);
}

func Init() {
	var err error;

	ctx, _ = context.WithTimeout(context.Background(), 10 * time.Second);
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"));
	err = client.Connect(ctx);

	check(err);

	database := client.Database("golang_marketplace");
	products = database.Collection("products");
}

func check(err error) {
	if (err != nil) {
		panic(err);
	}
}

func checkWithResponse(err error, responseWriter http.ResponseWriter) {
	if (err != nil) {
		fmt.Fprintf(responseWriter, err.Error());
		panic(err);
	}
}