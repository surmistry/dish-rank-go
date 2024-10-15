package handlers

import (
	"database/sql"
	"dish-rank-go/dish-rank-go/pkg/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	// "github.com/janirefdez/ArticleRestApi/pkg/models"
)

type handler struct {
	DB *sql.DB
}

func New(db *sql.DB) handler {
	return handler{db}
}

func (h handler) AddRestaurant(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
		w.WriteHeader(500)
		return
	}
	var restaurant models.Restaurant
	json.Unmarshal(body, &restaurant)

	restaurant.Id = (uuid.New()).String()
	queryStmt := `INSERT INTO restaurants (id,name,cuisine,address) VALUES ($1, $2, $3, $4) RETURNING id;`
	err = h.DB.QueryRow(queryStmt, &restaurant.Id, &restaurant.Name, &restaurant.Cuisine, &restaurant.Address).Scan(&restaurant.Id)
	if err != nil {
		log.Println("failed to execute query", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")

}

func (h handler) GetAllRestaurants(w http.ResponseWriter, r *http.Request) {

	results, err := h.DB.Query("SELECT * FROM restaurants;")
	if err != nil {
		log.Println("failed to execute query", err)
		w.WriteHeader(500)
		return
	}

	var restaurants = make([]models.Restaurant, 0)
	for results.Next() {
		var restaurant models.Restaurant
		err = results.Scan(&restaurant.Id, &restaurant.Name, &restaurant.Cuisine, &restaurant.Address)
		if err != nil {
			log.Println("failed to scan", err)
			w.WriteHeader(500)
			return
		}

		restaurants = append(restaurants, restaurant)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(restaurants)
}

func (h handler) GetRestaurant(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	queryStmt := `SELECT * FROM restaurants WHERE id = $1 ;`
	results, err := h.DB.Query(queryStmt, id)
	if err != nil {
		log.Println("failed to execute query", err)
		w.WriteHeader(500)
		return
	}

	var restaurant models.Restaurant
	for results.Next() {
		err = results.Scan(&restaurant.Id, &restaurant.Name, &restaurant.Cuisine, &restaurant.Address)
		if err != nil {
			log.Println("failed to scan", err)
			w.WriteHeader(500)
			return
		}
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(restaurant)
}

func (h handler) UpdateRestaurant(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Read request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedRestaurant models.Restaurant
	json.Unmarshal(body, &updatedRestaurant)

	queryStmt := `UPDATE restaurants SET name = $2, cuisine = $3, address = $4 WHERE id = $1 RETURNING id;`
	err = h.DB.QueryRow(queryStmt, &id, &updatedRestaurant.Name, &updatedRestaurant.Cuisine, &updatedRestaurant.Address).Scan(&id)
	if err != nil {
		log.Println("failed to execute query", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")

}

// func (h handler) DeleteRestaurant(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id := vars["id"]

// 	queryStmt := `DELETE FROM Restaurant WHERE id = $1;`
// 	_, err := h.DB.Query(queryStmt, &id)
// 	if err != nil {
// 		log.Println("failed to execute query", err)
// 		w.WriteHeader(500)
// 		return
// 	}

// 	w.Header().Add("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode("Deleted")
// }
