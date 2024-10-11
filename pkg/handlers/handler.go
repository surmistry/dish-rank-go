package handlers

import (
	"database/sql"
)

type handler struct {
	DB *sql.DB
}

func New(db *sql.DB) handler {
	return handler{db}
}

// func (h handler) GetAllArticles(w http.ResponseWriter, r *http.Request) {

// 	results, err := h.DB.Query("SELECT * FROM articles;")
// 	if err != nil {
// 		log.Println("failed to execute query", err)
// 		w.WriteHeader(500)
// 		return
// 	}

// 	var articles = make([]models.Article, 0)
// 	for results.Next() {
// 		var article models.Article
// 		err = results.Scan(&article.Id, &article.Title, &article.Desc, &article.Content)
// 		if err != nil {
// 			log.Println("failed to scan", err)
// 			w.WriteHeader(500)
// 			return
// 		}

// 		articles = append(articles, article)
// 	}

// 	w.Header().Add("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(articles)
// }

// func (h handler) GetArticle(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id := vars["id"]

// 	queryStmt := `SELECT * FROM articles WHERE id = $1 ;`
// 	results, err := h.DB.Query(queryStmt, id)
// 	if err != nil {
// 		log.Println("failed to execute query", err)
// 		w.WriteHeader(500)
// 		return
// 	}

// 	var article models.Article
// 	for results.Next() {
// 		err = results.Scan(&article.Id, &article.Title, &article.Desc, &article.Content)
// 		if err != nil {
// 			log.Println("failed to scan", err)
// 			w.WriteHeader(500)
// 			return
// 		}
// 	}

// 	w.Header().Add("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(article)
// }

// func (h handler) AddArticle(w http.ResponseWriter, r *http.Request) {
// 	// Read to request body
// 	defer r.Body.Close()
// 	body, err := ioutil.ReadAll(r.Body)

// 	if err != nil {
// 		log.Fatalln(err)
// 		w.WriteHeader(500)
// 		return
// 	}
// 	var article models.Article
// 	json.Unmarshal(body, &article)

// 	article.Id = (uuid.New()).String()
// 	queryStmt := `INSERT INTO articles (id,title,description,content) VALUES ($1, $2, $3, $4) RETURNING id;`
// 	err = h.DB.QueryRow(queryStmt, &article.Id, &article.Title, &article.Desc, &article.Content).Scan(&article.Id)
// 	if err != nil {
// 		log.Println("failed to execute query", err)
// 		w.WriteHeader(500)
// 		return
// 	}

// 	w.Header().Add("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode("Created")

// }

// func (h handler) UpdateArticle(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id := vars["id"]

// 	// Read request body
// 	defer r.Body.Close()
// 	body, err := ioutil.ReadAll(r.Body)

// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	var updatedArticle models.Article
// 	json.Unmarshal(body, &updatedArticle)

// 	queryStmt := `UPDATE articles SET title = $2, description = $3, content = $4 WHERE id = $1 RETURNING id;`
// 	err = h.DB.QueryRow(queryStmt, &id, &updatedArticle.Title, &updatedArticle.Desc, &updatedArticle.Content).Scan(&id)
// 	if err != nil {
// 		log.Println("failed to execute query", err)
// 		w.WriteHeader(500)
// 		return
// 	}

// 	w.Header().Add("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode("Updated")

// }

// func (h handler) DeleteArticle(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id := vars["id"]

// 	queryStmt := `DELETE FROM articles WHERE id = $1;`
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
