package main

import (
	"insta-api/controllers"
	"insta-api/databases"
	"log"
	"net/http"
)

func main() {
	//DB Connection
	databases.Connectdb()

	//Setting Routes
	mux := http.NewServeMux()
	mux.HandleFunc("/", controllers.Home)
	mux.HandleFunc("/users", controllers.UserHandler)
	mux.HandleFunc("/posts", controllers.PostHandler)
	mux.HandleFunc("/posts/users", controllers.BothHandler)

	//listening the Ports
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

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
