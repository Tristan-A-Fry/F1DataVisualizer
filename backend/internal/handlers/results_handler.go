package handlers

import (
  "encoding/json"
  "net/http"
  "F1DataVisualizer/backend/internal/services"
  "github.com/go-chi/chi/v5"
)

func ResultsDataHandler(w http.ResponseWriter, r *http.Request) {
    // Get the year parameter from the URL path
    year := chi.URLParam(r, "year")

    // Get the limit parameter from the query string, if provided
    limit := r.URL.Query().Get("limit")
    

    // Fetch race data for the specified year with an optional limit
    data, err := services.GetResultsData(year, limit)
    if err != nil {
        http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
        return
    }

    // Set Content-Type to application/json
    w.Header().Set("Content-Type", "application/json")

    // Encode data to JSON and send it as the HTTP response
    if err := json.NewEncoder(w).Encode(data); err != nil {
        http.Error(w, "Failed to encode data", http.StatusInternalServerError)
        return
    }
}

