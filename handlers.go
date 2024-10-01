package main

import (
    "encoding/json"
    "net/http"
)

// ShortenURLHandler handles requests to shorten URLs
func ShortenURLHandler(store *URLStore) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var reqData map[string]string
        if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
            http.Error(w, "Invalid request payload", http.StatusBadRequest)
            return
        }

        originalURL := reqData["url"]
        if originalURL == "" {
            http.Error(w, "URL cannot be empty", http.StatusBadRequest)
            return
        }

        shortURL := GenerateShortURL()
        store.SaveURL(shortURL, originalURL)

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{"short_url": shortURL})
    }
}

// RedirectURLHandler handles redirection from short URLs to original URLs
func RedirectURLHandler(store *URLStore) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        shortURL := r.URL.Path[len("/r/"):]

        originalURL, exists := store.GetOriginalURL(shortURL)
        if !exists {
            http.Error(w, "URL not found", http.StatusNotFound)
            return
        }

        http.Redirect(w, r, originalURL, http.StatusFound)
    }
}
