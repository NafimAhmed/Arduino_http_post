package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("hallow world")

	handleRoutes()

}

func homePage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Endpoint is hit")

}

func handleRoutes() {
	route := mux.NewRouter()

	route.HandleFunc("/", homePage).Methods("GET")

	http.ListenAndServe(":8080", route)
}
