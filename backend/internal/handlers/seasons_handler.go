package handlers

import(
  "encoding/json"
  "net/http"
  "f1_app/backend/internal/services"
)

func SeasonsDataHandler(w http.ResponseWriter, r *http.Request) {
    // Get the season query parameter from the URL
    season := r.URL.Query().Get("season")

    // Fetch data, passing the season if specified
    data, err := services.GetSeasonsData(season)
    if err != nil {
        http.Error(w, "Failed to fetch season data", http.StatusInternalServerError)
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

