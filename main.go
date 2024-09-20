package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var rankings []Ranking
var restaurants []Restaurant
var dishes []Dish

func greetingsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Greeting from Go Server 👋")
}

func getAllRestaurants(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(restaurants)
}

func addRestaurant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var restaurant Restaurant
	json.NewDecoder(r.Body).Decode(&restaurant)
	restaurants = append(restaurants, restaurant)

	json.NewEncoder(w).Encode(restaurants)
}

func getAllDishes(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(dishes)
}
func addDish(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var dish Dish
	json.NewDecoder(r.Body).Decode(&dish)
	dishes = append(dishes, dish)

	json.NewEncoder(w).Encode(dishes)
}

func main() {
	r := mux.NewRouter()
	var McDonalds = Restaurant{"McDonalds", "Junk"}
	restaurants = append(restaurants, McDonalds)
	dishes = append(dishes, Dish{"McNuggets", "Mixed meat packed into small shapes", McDonalds})
	r.HandleFunc("/", greetingsHandler).Methods("GET")
	r.HandleFunc("/restaurants", getAllRestaurants).Methods("GET")
	r.HandleFunc("/restaurants/add", addRestaurant).Methods("POST")
	r.HandleFunc("/dishes", getAllDishes).Methods("GET")
	r.HandleFunc("/dishes/add", addDish).Methods("POST")
	// r.HandleFunc("/add", addRanking).Methods("POST")

	fmt.Printf("Hello from GoServer 👋")
	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
