package main

import (
    "fmt"
    "log"
    "net/http"
    "F1DataVisualizer/backend/internal/handlers" 
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

func main() {
    
  r := chi.NewRouter()

  r.Use(middleware.Logger)
  r.Use(middleware.Recoverer)
  
  r.Get("/seasons", handlers.SeasonsDataHandler)

  // Start the HTTP server on localhost:8080
  fmt.Println("Server starting on http://localhost:8080")
  log.Fatal(http.ListenAndServe(":8080", r))
}


