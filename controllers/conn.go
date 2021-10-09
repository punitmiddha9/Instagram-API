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

func UserHandler(w http.ResponseWriter, r *http.Request) {
	usersCollection := databases.Mongo.DB.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	if r.Method == "GET" {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "The id query parameter is missing", http.StatusBadRequest)
			return
		}
		fmt.Fprint(w, id)
		filterCursor, err := usersCollection.Find(ctx, bson.M{"_id": id})
		if err != nil {
			log.Fatal(err)
		}
		var userFiltered []bson.M
		if err = filterCursor.All(ctx, &userFiltered); err != nil {
			log.Fatal(err)
		}
		fmt.Fprintln(w, userFiltered)
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {

}

func BothHandler(w http.ResponseWriter, r *http.Request) {

}
