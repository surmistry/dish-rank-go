package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/janirefdez/ArticleRestApi/pkg/models"
)

type handler struct {
	DB *sql.DB
}

func New(db *sql.DB) handler {
	return handler{db}
}

func (h handler) GetAllArticles(w http.ResponseWriter, r *http.Request) {

	results, err := h.DB.Query("SELECT * FROM articles;")
	if err != nil {
		log.Println("failed to execute query", err)
		w.WriteHeader(500)
		return
	}

	var articles = make([]models.Article, 0)
	for results.Next() {
		var article models.Article
		err = results.Scan(&article.Id, &article.Title, &article.Desc, &article.Content)
		if err != nil {
			log.Println("failed to scan", err)
			w.WriteHeader(500)
			return
		}

		articles = append(articles, article)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(articles)
}
