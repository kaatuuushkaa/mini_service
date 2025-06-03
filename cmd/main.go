package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/kaatuuushkaa/mini_service/internal/handler"
	"github.com/kaatuuushkaa/mini_service/internal/storage"
	"log"
	"net/http"
	"time"
)

var (
	db  *sql.DB
	err error
)

func main() {
	for i := 0; i < 10; i++ {
		db, err = storage.NewPostgresDB()
		if err == nil {
			break
		}
		log.Println("Waiting for database...")
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/quotes", handler.PostQuotesHandler(db)).Methods("POST")
	r.HandleFunc("/quotes", handler.GetQuotesHandler(db)).Methods("GET")
	r.HandleFunc("/quotes/random", handler.GetQuotesRandomHandler(db)).Methods("GET")
	r.HandleFunc("/quotes/{id}", handler.DeleteQuoteHandler(db)).Methods("DELETE")

	log.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", r)

	log.Println("Server stopped.")

}
