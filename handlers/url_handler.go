package handlers

import (
    "encoding/json"
    "fmt"
    "math/rand"
    "net/http"
    "ShortenUrlApp/models"
)

// Generate a random short code
func generateShortCode() string {
    const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    const length = 8
    shortCode := make([]byte, length)
    for i := range shortCode {
        shortCode[i] = charset[rand.Intn(len(charset))]
    }
    return string(shortCode)
}

// ShortenURL handles URL shortening requests
func ShortenURL(w http.ResponseWriter, r *http.Request) {
    var input struct {
        OriginalURL string `json:"original_url"`
    }
    
    w.Header().Add("Content-Type","application/json")
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }
    fmt.Println("input",input,r.Body)
    
    shortCode := generateShortCode()

    // Save the URL and short code to the database
    if err := models.SaveURL(input.OriginalURL, shortCode); err != nil {
        http.Error(w, "Failed to save URL", http.StatusInternalServerError)
        return
    }

    response := map[string]string{
        "short_url": fmt.Sprintf("http://localhost:8080/%s", shortCode),
    }
    json.NewEncoder(w).Encode(response)
}

// RedirectURL handles redirecting short codes to original URLs
func RedirectURL(w http.ResponseWriter, r *http.Request) {
    shortCode := r.URL.Path[1:]

    // Retrieve the original URL from the database
    originalURL, err := models.GetOriginalURL(shortCode)
    if err != nil {
        http.Error(w, "Short code not found", http.StatusNotFound)
        return
    }

    http.Redirect(w, r, originalURL, http.StatusFound)
}
