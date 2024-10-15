package handlers

import (
	"database/sql"
	// "github.com/janirefdez/ArticleRestApi/pkg/models"
)

type handler struct {
	DB *sql.DB
}

func New(db *sql.DB) handler {
	return handler{db}
}
