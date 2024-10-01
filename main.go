package main

import (
    "log"
    "net/http"
)

func main() {
    // Initialize the URL store
    store := NewURLStore()

    // Define HTTP handlers
    http.HandleFunc("/shorten", ShortenURLHandler(store))
    http.HandleFunc("/r/", RedirectURLHandler(store))

    // Start the HTTP server
    log.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
