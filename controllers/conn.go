package controllers

import (
	"context"
	"fmt"
	"insta-api/databases"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//Home Function
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello! This is the Home Page ") // send data to client side
}

//UserHandler Function
//"GET" and "POST" method for /users
func UserHandler(w http.ResponseWriter, r *http.Request) {
	usersCollection := databases.Mongo.DB.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	if r.Method == "GET" {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "The id query parameter is missing!", http.StatusBadRequest)
			return
		}
		filterCursor, err := usersCollection.Find(ctx, bson.M{"_id": id})
		if err != nil {
			log.Fatal(err)
		}
		var userFiltered []bson.M
		if err = filterCursor.All(ctx, &userFiltered); err != nil {
			log.Fatal(err)
		}
		fmt.Print(userFiltered)
	}
	if r.Method == "POST" {
		userResult, err := usersCollection.InsertOne(ctx, bson.D{
			{Key: "name", Value: "Punit Middha"},
			{Key: "email", Value: "punit.middha2019@vitstudent.ac.in"},
			{Key: "password", Value: "Punit@694"},
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Created a user!\n")
		fmt.Println(userResult)
	}
}

//PostHandler Function
//"GET" and "POST" method for /posts
func PostHandler(w http.ResponseWriter, r *http.Request) {
	postsCollection := databases.Mongo.DB.Collection("posts")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	//for GET request
	if r.Method == "GET" {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "The id query parameter is missing!", http.StatusBadRequest)
			return
		}
		filterCursor, err := postsCollection.Find(ctx, bson.M{"_id": id})
		if err != nil {
			log.Fatal(err)
		}
		var postFiltered []bson.M
		if err = filterCursor.All(ctx, &postFiltered); err != nil {
			log.Fatal(err)
		}
		fmt.Print(postFiltered)
	}

	//for POST request
	if r.Method == "POST" {
		postResult, err := postsCollection.InsertOne(ctx, bson.D{
			{Key: "image", Value: "https://www.bloorresearch.com/wp-content/uploads/2013/03/MONGO-DB-logo-300x470--x.png"},
			{Key: "caption", Value: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s."},
			{Key: "timestamp", Value: time.Now().String()},
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Created a post!\n")
		fmt.Println(postResult)
	}
}

//BothHandler Function
//"GET" method for /posts/users/{id}
func BothHandler(w http.ResponseWriter, r *http.Request) {
	postsCollection := databases.Mongo.DB.Collection("posts")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	//for GET request
	if r.Method == "GET" {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "The id query parameter is missing!", http.StatusBadRequest)
			return
		}
		filterCursor, err := postsCollection.Find(ctx, bson.M{"_id": id})
		if err != nil {
			log.Fatal(err)
		}
		var postFiltered []bson.M
		if err = filterCursor.All(ctx, &postFiltered); err != nil {
			log.Fatal(err)
		}
		fmt.Print(postFiltered)
	}
}
