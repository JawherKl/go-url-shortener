package main

import (
    "encoding/json"
    "log"
    "math/rand"
    "net/http"
    "sync"
    "time"
)

// URLStore is a thread-safe store for shortened URLs
type URLStore struct {
    sync.RWMutex
    URLs map[string]string // Correctly define the URLs map
}

// generateShortURL generates a random short string for the URL
func generateShortURL() string {
    rand.Seed(time.Now().UnixNano())
    letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
    b := make([]rune, 6)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

// shortenURL handles URL shortening requests
func (store *URLStore) shortenURL(w http.ResponseWriter, r *http.Request) {
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

    shortURL := generateShortURL()

    // Lock the store for writing
    store.Lock()
    store.URLs[shortURL] = originalURL // Store the original URL with the generated short URL
    store.Unlock()

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"short_url": shortURL})
}

// redirectURL handles redirection from a short URL to the original URL
func (store *URLStore) redirectURL(w http.ResponseWriter, r *http.Request) {
    shortURL := r.URL.Path[len("/r/"):]
    
    // Lock the store for reading
    store.RLock()
    originalURL, exists := store.URLs[shortURL] // Check if the short URL exists
    store.RUnlock()

    if !exists {
        http.Error(w, "URL not found", http.StatusNotFound)
        return
    }

    http.Redirect(w, r, originalURL, http.StatusFound)
}

func main() {
    // Initialize the URL store with an empty map
    store := &URLStore{
        URLs: make(map[string]string),
    }

    // Define HTTP handlers
    http.HandleFunc("/shorten", store.shortenURL)
    http.HandleFunc("/r/", store.redirectURL)

    log.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
