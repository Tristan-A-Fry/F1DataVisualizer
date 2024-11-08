package services

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
    "F1DataVisualizer/backend/config"
)

func GetDriverStandingsData(year string) (interface{}, error) {
    // Construct the API URL with the year as part of the path
    apiURL := fmt.Sprintf("%s/%s/driverstandings", config.APIBaseURL, year)

    // Log the final URL for debugging
    fmt.Println("Final URL:", apiURL)

    // Create the request
    req, err := http.NewRequest("GET", apiURL, nil)
    if err != nil {
        return nil, fmt.Errorf("failed to create request: %w", err)
    }

    client := &http.Client{Timeout: 10 * time.Second}
    resp, err := client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("failed to make request: %w", err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("failed to read response: %w", err)
    }

    var data interface{}
    if err := json.Unmarshal(body, &data); err != nil {
        return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
    }

    return data, nil
}
