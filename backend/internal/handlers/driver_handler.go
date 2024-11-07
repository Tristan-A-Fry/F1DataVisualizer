package handlers

import(
  "encoding/json"
  "net/http"
  "f1_app/backend/internal/services"
)

func DriverDataHandler(w http.ResponseWriter, r *http.Request) {
    // Retrieve query parameters from the request
    driverNumber := r.URL.Query().Get("driver_number")
    sessionKey := r.URL.Query().Get("session_key")

    // Check if required parameters are provided
    if driverNumber == "" || sessionKey == "" {
        http.Error(w, "Missing required query parameters: driver_number and session_key", http.StatusBadRequest)
        return
    }

    // Fetch data from the API
    data, err := services.GetDriverData(driverNumber, sessionKey)
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
