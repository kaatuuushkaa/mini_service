package storage

import (
	"database/sql"
	"fmt"
	"github.com/kaatuuushkaa/mini_service/internal/models"
	_ "github.com/lib/pq"
	"math/rand"
)

func PostQuotes(db *sql.DB, quote, author string) error {
	var id int

	err := db.QueryRow("INSERT INTO quotes(author, quote) VALUES ($1, $2) RETURNING id", author, quote).Scan(&id)
	if err != nil {
		return fmt.Errorf("Error creating quote: %w", err)
	}

	return nil
}

func GetQuotes(db *sql.DB) ([]*models.Quotes, error) {
	var quotes []*models.Quotes

	rows, err := db.Query("SELECT id,author, quote FROM quotes")
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("No quotes found")
		}
		return nil, fmt.Errorf("Could not get quotes: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var q models.Quotes
		err := rows.Scan(&q.ID, &q.Author, &q.Quote)
		if err != nil {
			return nil, fmt.Errorf("Could not get quotes: %v", err)
		}
		quotes = append(quotes, &q)
	}

	return quotes, nil
}

func GetQuoteRandom(db *sql.DB) (*models.Quotes, error) {
	var quotes models.Quotes
	var totalRows int

	err := db.QueryRow("SELECT COUNT(*) AS total_rows FROM quotes").Scan(&totalRows)
	if err != nil {
		return nil, fmt.Errorf("Could not count rows: %v", err)
	}

	if totalRows == 0 {
		return nil, fmt.Errorf("No quotes found")
	}

	offset := rand.Intn(totalRows)

	err = db.QueryRow("SELECT id, author, quote FROM quotes OFFSET $1 LIMIT 1", offset).
		Scan(&quotes.ID, &quotes.Author, &quotes.Quote)
	if err != nil {
		return nil, fmt.Errorf("Could not get quote: %v", err)
	}

	return &quotes, nil
}

func GetQuotesAuthor(db *sql.DB, author string) ([]*models.Quotes, error) {
	var quotes []*models.Quotes

	rows, err := db.Query("SELECT id, author, quote FROM quotes WHERE author = $1", author)
	if err != nil {
		return nil, fmt.Errorf("Fail to query database")
	}
	defer rows.Close()

	for rows.Next() {
		var q models.Quotes
		err := rows.Scan(&q.ID, &q.Author, &q.Quote)
		if err != nil {
			return nil, fmt.Errorf("Failed to scan row")
		}
		quotes = append(quotes, &q)
	}

	return quotes, nil
}

func DeleteQuote(db *sql.DB, quoteID int) error {
	_, err := db.Exec("DELETE FROM quotes WHERE id = $1", quoteID)
	if err != nil {
		return fmt.Errorf("Error deleting quote: %v", err)
	}
	return nil
}
