package main

import (
    "math/rand"
    "sync"
    "time"
)

type URLStore struct {
    sync.RWMutex
    URLs map[string]string
}

// NewURLStore initializes and returns a new URLStore instance
func NewURLStore() *URLStore {
    return &URLStore{
        URLs: make(map[string]string),
    }
}

// GenerateShortURL generates a random 6-character string for the short URL
func GenerateShortURL() string {
    rand.Seed(time.Now().UnixNano())
    letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
    b := make([]rune, 6)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

// SaveURL stores the mapping of a short URL to the original URL
func (store *URLStore) SaveURL(shortURL, originalURL string) {
    store.Lock()
    store.URLs[shortURL] = originalURL
    store.Unlock()
}

// GetOriginalURL retrieves the original URL for a given short URL
func (store *URLStore) GetOriginalURL(shortURL string) (string, bool) {
    store.RLock()
    originalURL, exists := store.URLs[shortURL]
    store.RUnlock()
    return originalURL, exists
}
