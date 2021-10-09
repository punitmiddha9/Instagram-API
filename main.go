package main

import (
	"fmt"
	"insta-api/databases"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello! This is the Home Page ") // send data to client side
}

func main() {
	databases.Connectdb()

	http.HandleFunc("/", Home)                 // set router
	error := http.ListenAndServe(":9090", nil) // set listen port
	if error != nil {
		log.Fatal("ListenAndServe: ", error)
	}

}
