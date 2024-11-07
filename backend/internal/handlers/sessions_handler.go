
package handlers

import(
  "encoding/json"
  "net/http"
  "f1_app/backend/internal/services"
)

func SessionsDataHandler(w http.ResponseWriter, r *http.Request) {
    // Retrieve query parameters from the request
    countryName := r.URL.Query().Get("country_name")
    sessionName := r.URL.Query().Get("session_name")
    nYear := r.URL.Query().Get("year")

    // Check if required parameters are provided
    if countryName == "" || sessionName == "" || nYear == "" {
        http.Error(w, "Missing required query parameters: country_name and session_name and year", http.StatusBadRequest)
        return
    }

    // Fetch data from the API
    data, err := services.GetSessionsData(countryName, sessionName, nYear)
    if err != nil {
        http.Error(w, "Failed to fetch driver data", http.StatusInternalServerError)
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
