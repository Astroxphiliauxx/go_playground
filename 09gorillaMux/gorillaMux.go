package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Welcome to the Gorilla Mux example!")

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")       //HomeHandler is a func, it is here used for reference
	r.HandleFunc("/about", AboutHandler).Methods("GET") //first para is the path

	//to call the errors, we use log fatal, and pass the http.ListenAndServe function
	fmt.Println("Starting server on :8080")

	log.Fatal(http.ListenAndServe(":8080", r)) //r is the router we created

}

//these are the funct, which take two params, w = writer and r = request, request is a pointer

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Welcome to the Home Page!"))
	fmt.Fprintf(w, "Welcome to the Home Page!")

}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the About Page!")
}
