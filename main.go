package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var client *mongo.Client

//Declare global variable to identify posts of specific users
var u_id primitive.ObjectID

//New user structure
type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
}

//New post structure
type PostPic struct {
	ID      primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Posted  primitive.Timestamp `json:"posted,omitempty" bson:"posted,omitempty"`
	Userid  primitive.ObjectID  `json:"userid,omitempty" bson:"userid,omitempty"`
	Caption string              `json:"caption,omitempty" bson:"caption,omitempty"`
	Imgurl  string              `json:"imgurl,omitempty" bson:"imgurl,omitempty"`
}

//Convert string password to hash value
func getHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		fmt.Println("Password conversion error : ", err)
	}
	return string(hash)
}

//Create a new user
func CreateUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")

	var user User
	_ = json.NewDecoder(request.Body).Decode(&user)

	//Securely store password
	user.Password = getHash([]byte(user.Password))

	//DB name : instagram, collection name : users
	collection := client.Database("instagram").Collection("users")

	//Update current user id
	u_id = user.ID

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, user)
	json.NewEncoder(response).Encode(result)
}

//Get details of one user
func GetOneUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var user User
	collection := client.Database("instagram").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := collection.FindOne(ctx, User{ID: id}).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(user)
}

//Create post
func CreatePost(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")

	var post PostPic
	_ = json.NewDecoder(request.Body).Decode(&post)

	post.Userid = u_id

	//DB name : instagram, collection name : posts
	collection := client.Database("instagram").Collection("posts")

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, post)
	json.NewEncoder(response).Encode(result)
}

//Get a post using id
func GetOnePost(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var post PostPic
	collection := client.Database("instagram").Collection("posts")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := collection.FindOne(ctx, PostPic{ID: id}).Decode(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(post)
}

//List all posts of a user using userid
func GetAllPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")

	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var posts []PostPic
	collection := client.Database("instagram").Collection("post")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, PostPic{Userid: id})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var pic PostPic
		cursor.Decode(&pic)
		posts = append(posts, pic)

	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(posts)
}

func main() {
	fmt.Println("App started...")

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)

	//Define router
	router := mux.NewRouter()

	//Create a user
	router.HandleFunc("/users", CreateUser).Methods("POST")

	//Get a user using id
	router.HandleFunc("/users/{id}", GetOneUser).Methods("GET")

	//Create a post
	router.HandleFunc("/posts", CreatePost).Methods("POST")

	//Get post using id
	router.HandleFunc("/posts/{id}", GetOnePost).Methods("GET")

	//List all posts of a user
	router.HandleFunc("/posts/users/{id}", GetAllPosts).Methods("GET")

	//Start http server on port 9090
	http.ListenAndServe(":9090", router)
}
