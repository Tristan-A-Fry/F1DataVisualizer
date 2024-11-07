package handlers

import(
  "encoding/json" //allows us to decode and encode json data i.e convert Go data structures to json
  "net/http" //http handling
  "F1DataVisualizer/backend/internal/services"
)

func CircuitsDataHandler(w http.ResponseWriter, r *http.Request) {
    data, err := services.GetCircuitsData()
    if err != nil {
        http.Error(w, "Failed to fetch circuits data", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")

    if err := json.NewEncoder(w).Encode(data); err != nil {
        http.Error(w, "Failed to encode data", http.StatusInternalServerError)
        return
    }
}
