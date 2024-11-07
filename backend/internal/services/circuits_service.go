
package services

import (
    "encoding/json"
    "fmt" //format for I/O i.e strings and printed logs
    "io/ioutil" //used for reading data from an io.Reader to read the response body
    "net/http"
    "net/url" //used to parse and manipulate URLs
    "time"
    "F1DataVisualizer/backend/config" //geting the APIBaseURL from the config file 
)

func GetCircuitsData() (interface{}, error) {
    apiURL, err := url.Parse(fmt.Sprintf("%s/circuits", config.APIBaseURL))
    if err != nil {
        return nil, fmt.Errorf("failed to parse base URL: %w", err)
    }

    query := apiURL.Query()
    query.Set("limit", "77")
    apiURL.RawQuery = query.Encode()

    fmt.Println("Final URL:", apiURL.String())

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


