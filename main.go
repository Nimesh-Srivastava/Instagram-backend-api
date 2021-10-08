package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
}
