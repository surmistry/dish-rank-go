package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Restaurant struct {
	Name    string
	Cuisine string
	// Address	Address
}

type Dish struct {
	Name        string
	Description string
}

type Review struct {
	Comment string
	Dish    Dish
	// User		User
	// Asset	Asset
}

type Ranking struct {
	Previous Review
	Next     Review
	Review   Review
}

var rankings []Ranking

func greetingsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Greeting from Go Server 👋")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", greetingsHandler).Methods("GET")
	// r.HandleFunc("/add", addRanking).Methods("POST")

	fmt.Printf("Hello from GoServer 👋")
	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
