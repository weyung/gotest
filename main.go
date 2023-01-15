// main.go
package main

import (
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/weyung/gotest/config"
	"github.com/weyung/gotest/handler"
)

var db *sqlx.DB

func main() {
	var err error
	db, err = sqlx.Connect("postgres", config.GetDSN())
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	http.HandleFunc("/query", handler.QueryScore)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
