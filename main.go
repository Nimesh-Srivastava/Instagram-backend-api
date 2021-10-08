package main

import (
	"context"
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

func main() {
	fmt.Println("Starting the app...")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	//Set router
	http.HandleFunc("/", nil)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
