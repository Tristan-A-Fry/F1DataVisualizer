
package handlers

import (
  "encoding/json"
  "net/http"
  "F1DataVisualizer/backend/internal/services"
  "github.com/go-chi/chi/v5"
)

func ConstructorStandingsDataHandler(w http.ResponseWriter, r *http.Request) {
    // Get the year parameter from the URL path
    year := chi.URLParam(r, "year")

    // Fetch race data for the specified year
    data, err := services.GetConstructorStandingsData(year)
    if err != nil {
        http.Error(w, "Failed to fetch constructor standings data", http.StatusInternalServerError)
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

