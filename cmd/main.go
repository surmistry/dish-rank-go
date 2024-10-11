package main

import (
	"database/sql"
	"dish-rank-go/dish-rank-go/pkg/db"
	"dish-rank-go/dish-rank-go/pkg/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests(DB *sql.DB) {
	h := handlers.New(DB)
	// 	// create a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// 	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", h.GetAllArticles).Methods(http.MethodGet)
	myRouter.HandleFunc("/articles/{id}", h.GetArticle).Methods(http.MethodGet)
	myRouter.HandleFunc("/articles", h.AddArticle).Methods(http.MethodPost)
	myRouter.HandleFunc("/restaurants", h.Addrestaurant).Methods(http.MethodPost)
	myRouter.HandleFunc("/articles/{id}", h.UpdateArticle).Methods(http.MethodPut)
	myRouter.HandleFunc("/articles/{id}", h.DeleteArticle).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	DB := db.Connect()
	db.CreateTable(DB)
	handleRequests(DB)
	db.CloseConnection(DB)
}
