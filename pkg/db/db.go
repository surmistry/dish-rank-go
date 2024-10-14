package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/janirefdez/ArticleRestApi/pkg/mocks"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "surmistry"
	password = "password"
	dbname   = "dishRankDB"
)

func Connect() *sql.DB {
	connInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connInfo)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to db!")
	return db
}

func CloseConnection(db *sql.DB) {
	defer db.Close()
}

func CreateTable(db *sql.DB) {
	var exists bool
	if err := db.QueryRow("SELECT EXISTS (SELECT FROM pg_tables WHERE  schemaname = 'public' AND tablename = 'restaurants' );").Scan(&exists); err != nil {
		fmt.Println("failed to execute query", err)
		return
	}
	if !exists {
		results, err := db.Query("CREATE TABLE restaurants (id VARCHAR(36) PRIMARY KEY, title VARCHAR(100) NOT NULL, cuisinel VARCHAR(50) NOT NULL, address VARCHAR(50) NOT NULL);")
		if err != nil {
			fmt.Println("failed to execute query", err)
			return
		}
		fmt.Println("Table created successfully", results)

		for _, article := range mocks.Restaurants {
			queryStmt := `INSERT INTO restaurants (id,title,cuisine,address) VALUES ($1, $2, $3, $4) RETURNING id;`

			err := db.QueryRow(queryStmt, &article.Id, &article.Title, &article.Desc, &article.Content).Scan(&article.Id)
			if err != nil {
				log.Println("failed to execute query", err)
				return
			}
		}
		fmt.Println("Mock Articles included in Table", results)
	} else {
		fmt.Println("Table 'articles' already exists ")
	}

}
