package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"local.com/golangMarketplace/productService/models"
)

var ctx context.Context
var client *mongo.Client
var products *mongo.Collection

func Create(responseWriter http.ResponseWriter, request *http.Request) {
	var product models.AlterProduct

	if checkErrorWithResponse(json.NewDecoder(request.Body).Decode(&product), responseWriter) == true {
		return
	}

	if product.JwtEmail == "" {
		responseWriter.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(responseWriter, "Email is missing!")
		return
	}

	if product.Name == "" || product.Description == "" || product.Price == 0 {
		responseWriter.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(responseWriter, "Invalid product properties! Name: %s, description: %s, price: %d\n", product.Name, product.Description, product.Price)
		return
	}

	createdProduct, err := products.InsertOne(ctx, bson.D{
		{Key: "email", Value: product.JwtEmail},
		{Key: "name", Value: product.Name},
		{Key: "description", Value: product.Description},
		{Key: "price", Value: product.Price},
	})

	if checkErrorWithResponse(err, responseWriter) == true {
		return
	}

	var newObjectId primitive.ObjectID = createdProduct.InsertedID.(primitive.ObjectID)

	fmt.Fprintf(responseWriter, newObjectId.Hex())

}

func GetAll(responseWriter http.ResponseWriter, request *http.Request) {
	var foundProducts []models.Product
	cursor, err := products.Find(ctx, bson.M{})
	if checkErrorWithResponse(err, responseWriter) == true {
		return
	}

	for cursor.Next(ctx) {
		var foundProduct models.Product

		if checkErrorWithResponse(cursor.Decode(&foundProduct), responseWriter) == true {
			return
		}
		foundProducts = append(foundProducts, foundProduct)
	}

	jsonResponse, _ := json.Marshal(foundProducts)
	fmt.Fprintf(responseWriter, string(jsonResponse))
}

func GetById(responseWriter http.ResponseWriter, request *http.Request) {
	var foundProduct models.Product
	var id models.ObjectID

	check(json.NewDecoder(request.Body).Decode(&id))

	objectId, err := primitive.ObjectIDFromHex(id.Id)

	if checkErrorWithResponse(err, responseWriter) == true {
		return
	}

	err = products.FindOne(ctx, bson.M{"_id": objectId}).Decode(&foundProduct)

	if checkErrorWithResponse(err, responseWriter) == true {
		return
	}

	jsonResponse, _ := json.Marshal(foundProduct)
	fmt.Fprintf(responseWriter, string(jsonResponse))
}

func GetByEmail(responseWriter http.ResponseWriter, request *http.Request) {
	var foundProducts []models.Product
	var urlTokens []string = strings.Split(request.URL.Path, "/")
	var sellerEmail string = urlTokens[len(urlTokens)-1]

	cursor, err := products.Find(ctx, bson.M{"email": sellerEmail})

	for cursor.Next(ctx) {
		var foundProduct models.Product

		if checkErrorWithResponse(cursor.Decode(&foundProduct), responseWriter) == true {
			return
		}
		foundProducts = append(foundProducts, foundProduct)
	}

	if checkErrorWithResponse(err, responseWriter) == true {
		return
	}

	jsonResponse, _ := json.Marshal(foundProducts)
	fmt.Fprintf(responseWriter, string(jsonResponse))

}

func Update(responseWriter http.ResponseWriter, request *http.Request) {
	var newProductData models.AlterProduct
	var foundProduct models.Product

	if checkErrorWithResponse(json.NewDecoder(request.Body).Decode(&newProductData), responseWriter) == true {
		return
	}

	objectId, err := primitive.ObjectIDFromHex(newProductData.Id)
	if checkErrorWithResponse(err, responseWriter) == true {
		return
	}

	err = products.FindOne(ctx, bson.M{"_id": objectId}).Decode(&foundProduct)
	if checkErrorWithResponse(err, responseWriter) == true {
		return
	}

	if foundProduct.Email != newProductData.JwtEmail {
		responseWriter.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(responseWriter, "You can't update a product with someone else's email!")
		return
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
		})

	jsonResponse, _ := json.Marshal(updatedObject)
	fmt.Fprintf(responseWriter, string(jsonResponse))
}

func Delete(responseWriter http.ResponseWriter, request *http.Request) {
	var id models.ObjectID
	var result *mongo.DeleteResult
	var foundProduct models.AlterProduct

	if checkErrorWithResponse(json.NewDecoder(request.Body).Decode(&id), responseWriter) == true {
		return
	}

	objectId, err := primitive.ObjectIDFromHex(id.Id)
	if checkErrorWithResponse(err, responseWriter) == true {
		return
	}

	err = products.FindOne(ctx, bson.M{"_id": objectId}).Decode(&foundProduct)
	if checkErrorWithResponse(err, responseWriter) == true {
		return
	}

	if foundProduct.Email != id.JwtEmail {
		responseWriter.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(responseWriter, "You can't delete a product with someone else's email!")
		return
	}

	result, err = products.DeleteMany(ctx, bson.M{"_id": objectId})

	if checkErrorWithResponse(err, responseWriter) == true {
		return
	}

	fmt.Fprintf(responseWriter, "Deleted documents: %d", result.DeletedCount)
}

func init() {
	var err error

	ctx = context.Background()
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://admin:password@192.168.1.13:27017"))
	err = client.Connect(ctx)

	check(err)

	database := client.Database("golang_marketplace")
	products = database.Collection("products")
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func checkErrorWithResponse(err error, responseWriter http.ResponseWriter) bool {
	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(responseWriter, err.Error())
		fmt.Println(err.Error())
		return true
	}
	return false
}
