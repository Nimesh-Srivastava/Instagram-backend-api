package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Id    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name,omitempty" bson:"name,omitempty"`
	Email string             `json:"email,omitempty" bson:"email,omitempty"`
	Pswd  string             `json:"pswd,omitempty" bson:"pswd,omitempty"`
}

var client *mongo.Client

func createUser(response http.ResponseWriter, request *http.Request) {

	response.Header().Add("content-type", "application/json")

	var U User
	json.NewDecoder(request.Body).Decode(&U)

	collection := client.Database("instagram").Collection("users")
	ctx, err_ctx := context.WithTimeout(context.Background(), 10*time.Second)
	if err_ctx != nil {
		log.Fatal("ListenAndServe: ", err_ctx)
	}
	result, _ := collection.InsertOne(ctx, U)
	json.NewEncoder(response).Encode(result)
}

func main() {
	fmt.Println("App started...")

	ctx, err_ctx := context.WithTimeout(context.Background(), 10*time.Second)
	if err_ctx != nil {
		log.Fatal("ListenAndServe: ", err_ctx)
	}
	// client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	//Set router
	router := mux.NewRouter()

	//Create a user
	router.HandleFunc("/user", createUser).Methods("POST")

	err := http.ListenAndServe(":1234", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
