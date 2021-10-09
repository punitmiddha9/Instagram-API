package routes

import (
	"insta-api/controllers"
	"net/http"
)

func Route() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", controllers.UserHandler)
	mux.HandleFunc("/posts", controllers.PostHandler)
	mux.HandleFunc("/posts/users", controllers.BothHandler)
}

//Another Method to make the routes and related function without controllers folder
// func userhandler(w http.ResponseWriter, r *http.Request) {

// }

// func posthandler(w http.ResponseWriter, r *http.Request) {

// }

// func bothhandler(w http.ResponseWriter, r *http.Request) {

// }
// func Route()  {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/users", userhandler)
// 	mux.HandleFunc("/posts", posthandler)
// 	mux.HandleFunc("/posts/users", bothhandler)
// }
