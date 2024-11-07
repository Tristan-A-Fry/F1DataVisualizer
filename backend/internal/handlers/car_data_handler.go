package handlers

import(
  "net/http"
  "encoding/json"
  "f1_app/backend/internal/services"
)

// CarDataHandler handles requests to /car_data
func CarDataHandler(w http.ResponseWriter, r *http.Request) {
    // Define query parameters
    driverNumber := r.URL.Query().Get("driver_number")
    sessionKey := r.URL.Query().Get("session_key")
    speed := r.URL.Query().Get("speed")

    // Fetch data from the API
    data, err := services.GetCarData(driverNumber, sessionKey, speed)
    if err != nil {
        http.Error(w, "Failed to fetch car data", http.StatusInternalServerError)
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
