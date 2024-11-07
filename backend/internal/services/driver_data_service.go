package services

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
    "time"
    "f1_app/backend/config"
)

func GetDriverData(driverNumber, sessionKey string) (interface{}, error) {

    apiURL, err := url.Parse(fmt.Sprintf("%s/drivers", config.APIBaseURL))
    if err != nil {
        return nil, fmt.Errorf("failed to parse base URL: %w", err)
    }

    // Add query parameters
    query := apiURL.Query()
    query.Set("driver_number", driverNumber)
    query.Set("session_key", sessionKey)
    apiURL.RawQuery = query.Encode()

    // Log the final URL for debugging
    fmt.Println("Final URL:", apiURL.String())

    // Create the request
    req, err := http.NewRequest("GET", apiURL.String(), nil)
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

