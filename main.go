package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	id    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	name  string             `json:"name,omitempty" bson:"name,omitempty"`
	email string             `json:"email,omitempty" bson:"email,omitempty"`
	pswd  string             `json:"pswd,omitempty" bson:"pswd,omitempty"`
}

var client *mongo.Client

// func initialize(w http.ResponseWriter, r *http.Request) {}

func createUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	var user User
	json.NewDecoder(request.Body).Decode(&user)

	collection := client.Database("instagram").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, user)
	json.NewEncoder(response).Encode(result)
}

func main() {
	fmt.Println("Starting the app...")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	//Set router
	http.HandleFunc("/user", createUser)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
