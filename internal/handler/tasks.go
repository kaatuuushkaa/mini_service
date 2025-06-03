package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kaatuuushkaa/mini_service/internal/models"
	"github.com/kaatuuushkaa/mini_service/internal/storage"
	"net/http"
	"strconv"
)

func PostQuotesHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Quote  string `json:"quote"`
			Author string `json:"author"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
			return
		}

		err := storage.PostQuotes(db, req.Quote, req.Author)
		if err != nil {
			http.Error(w, fmt.Sprintf("Fail to create quote: %v", err), http.StatusInternalServerError)
			return
		}
	}
}

func GetQuotesHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var quotes []*models.Quotes
		author := r.URL.Query().Get("author")
		if author != "" {
			quotes, err = storage.GetQuotesAuthor(db, author)
		} else {
			quotes, err = storage.GetQuotes(db)
		}

		if err != nil {
			http.Error(w, fmt.Sprintf("Fail to get quotes: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(quotes)
	}
}

func GetQuotesRandomHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		quotes, err := storage.GetQuoteRandom(db)
		if err != nil {
			http.Error(w, fmt.Sprintf("Fail to get quotes: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(quotes)
	}
}

func DeleteQuoteHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		quoteIDStr := vars["id"]

		quoteID, err := strconv.Atoi(quoteIDStr)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid user ID: %v", err), http.StatusBadRequest)
			return
		}

		err = storage.DeleteQuote(db, quoteID)
		if err != nil {
			http.Error(w, "Fail to delete quote", http.StatusInternalServerError)
			return
		}
	}
}
