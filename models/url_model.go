package models

import (
    "database/sql"
    "fmt"
    "ShortenUrlApp/database"
)

type URL struct {
    ID         int64
    OriginalURL string
    ShortCode  string
}

// SaveURL stores the original URL and short code in the database
func SaveURL(originalURL, shortCode string) error {
    query := "INSERT INTO url_shortener (original_url, short_code) VALUES (:1, :2)"
    fmt.Println(originalURL,shortCode)
    _, err := database.DB.Exec(query, originalURL, shortCode)
    if err != nil {
        return fmt.Errorf("could not insert URL: %v", err)
    }
    return nil
}

// GetOriginalURL retrieves the original URL using the short code
func GetOriginalURL(shortCode string) (string, error) {
    var originalURL string
    query := "SELECT original_url FROM url_shortener WHERE short_code = :1"
    err := database.DB.QueryRow(query, shortCode).Scan(&originalURL)
    if err == sql.ErrNoRows {
        return "", fmt.Errorf("short code not found")
    } else if err != nil {
        return "", fmt.Errorf("error querying URL: %v", err)
    }
    return originalURL, nil
}
